//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const (
	baseURL = "http://127.0.0.1:8888/api"
)

func main() {
	// 获取两个测试 Token
	tokenA := getToken("Bradtke1974284")
	tokenB := getToken("DuBuque1671714")

	if tokenA == "" || tokenB == "" {
		log.Fatal("Failed to get tokens")
	}
	fmt.Println("Tokens retrieved successfully.")

	// 测试 1: 查询帖子列表，获取一个测试用 PostID
	postID := getFirstPostID(tokenA)
	if postID == 0 {
		log.Fatal("No posts found to test with.")
	}
	fmt.Printf("Using Post ID: %d for testing\n", postID)

	// 测试 2: 用户A 对帖子点赞
	fmt.Println("\n--- Test 1: User A likes the post ---")
	likeResA1 := toggleLikePost(tokenA, postID)
	if !likeResA1.IsLiked {
		log.Fatal("Expected IsLiked to be true on first like, got false")
	}
	fmt.Println("SUCCESS: User A liked the post. IsLiked =", likeResA1.IsLiked)

	// 测试 3: 用户B 查看帖子详情，IsLiked 应该为 false
	fmt.Println("\n--- Test 2: User B views the post (should not be liked) ---")
	viewResB1 := viewPost(tokenB, postID)
	if viewResB1.IsLiked {
		log.Fatal("Expected User B to see IsLiked=false, but got true")
	}
	fmt.Println("SUCCESS: User B sees IsLiked =", viewResB1.IsLiked)

	// 测试 4: 用户A 再次对帖子点赞 (取消点赞)
	fmt.Println("\n--- Test 3: User A unlikes the post ---")
	likeResA2 := toggleLikePost(tokenA, postID)
	if likeResA2.IsLiked {
		log.Fatal("Expected IsLiked to be false on second like (unlike), got true")
	}
	fmt.Println("SUCCESS: User A unliked the post. IsLiked =", likeResA2.IsLiked)

	// 测试 5: 用户A 查看帖子详情，IsLiked 应该为 false
	fmt.Println("\n--- Test 4: User A views the post after unlike (should not be liked) ---")
	viewResA1 := viewPost(tokenA, postID)
	if viewResA1.IsLiked {
		log.Fatal("Expected User A to see IsLiked=false after unliking, but got true")
	}
	fmt.Println("SUCCESS: User A sees IsLiked =", viewResA1.IsLiked)

	fmt.Println("\n==============================================")
	fmt.Println("TESTING COMMENT LIKES")
	fmt.Println("==============================================")
	
	// 测试 6: 用户A 发表一条测试评论，为了测试点赞
	commentID := createComment(tokenA, postID, "This is a test comment for liking")
	if commentID == 0 {
		log.Fatal("Failed to create test comment")
	}
	fmt.Printf("Created Test Comment ID: %d\n", commentID)
	
	// 测试 7: 用户A 对评论点赞
	fmt.Println("\n--- Test 5: User A likes the comment ---")
	likeCmtResA1 := toggleLikeComment(tokenA, commentID)
	if !likeCmtResA1.IsLiked {
		log.Fatal("Expected comment IsLiked to be true on first like, got false")
	}
	fmt.Println("SUCCESS: User A liked the comment. IsLiked =", likeCmtResA1.IsLiked)

	// 测试 8: 用户A 再次对评论点赞 (取消)
	fmt.Println("\n--- Test 6: User A unlikes the comment ---")
	likeCmtResA2 := toggleLikeComment(tokenA, commentID)
	if likeCmtResA2.IsLiked {
		log.Fatal("Expected comment IsLiked to be false on unlike, got true")
	}
	fmt.Println("SUCCESS: User A unliked the comment. IsLiked =", likeCmtResA2.IsLiked)

	fmt.Println("\n==============================================")
	fmt.Println("ALL INTEGRATION TESTS PASSED SUCCESSFULLY!")
	fmt.Println("==============================================")
}

type commentCreateRes struct {
	Id uint64 `json:"id"`
}

func createComment(token string, postID uint64, content string) uint64 {
	reqBody := map[string]interface{}{
		"postId": postID,
		"content": content,
		"userId": 1, // 会被中间件的用户ID覆盖
	}
	var res commentCreateRes
	err := doReq("POST", "/forum/comments/create", token, reqBody, &res)
	if err != nil {
		log.Fatal("createComment err:", err)
	}
	return res.Id
}

func toggleLikeComment(token string, commentID uint64) likeRes {
	reqBody := map[string]interface{}{"id": commentID}
	var res likeRes
	err := doReq("POST", "/forum/comments/like", token, reqBody, &res)
	if err != nil {
		log.Fatal("toggleLikeComment err:", err)
	}
	return res
}

func getToken(account string) string {
	cmd := exec.Command("go", "run", "scripts/get_token.go")
	cmd.Env = append(cmd.Environ(), "ACCOUNT="+account)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run get_token.go for %s: %v", account, err)
		return ""
	}
	return out.String()[:len(out.String())-1] // remove trailing newline
}

func doReq(method, path, token string, body interface{}, resObj interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(b)
	}
	req, _ := http.NewRequest(method, baseURL+path, bodyReader)
	req.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	resBody, _ := io.ReadAll(resp.Body)
	
	// GF default response wrapper
	var wrapper struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(resBody, &wrapper); err != nil {
		return err
	}
	if wrapper.Code != 0 {
		return fmt.Errorf("API Error: %s", wrapper.Message)
	}

	if resObj != nil && len(wrapper.Data) > 0 {
		return json.Unmarshal(wrapper.Data, resObj)
	}
	return nil
}

type listRes struct {
	List []struct {
		Id uint64 `json:"id"`
	} `json:"list"`
}

func getFirstPostID(token string) uint64 {
	var res listRes
	err := doReq("GET", "/forum/posts/list", token, nil, &res)
	if err != nil {
		log.Fatal("getFirstPostID err:", err)
	}
	if len(res.List) > 0 {
		return res.List[0].Id
	}
	return 0
}

type likeRes struct {
	IsLiked bool `json:"isLiked"`
}

func toggleLikePost(token string, postID uint64) likeRes {
	reqBody := map[string]interface{}{"id": postID}
	var res likeRes
	err := doReq("POST", "/forum/posts/like", token, reqBody, &res)
	if err != nil {
		log.Fatal("toggleLikePost err:", err)
	}
	return res
}

type viewRes struct {
	IsLiked bool `json:"isLiked"`
}

func viewPost(token string, postID uint64) viewRes {
	var res viewRes
	err := doReq("GET", fmt.Sprintf("/forum/posts/view?id=%d", postID), token, nil, &res)
	if err != nil {
		log.Fatal("viewPost err:", err)
	}
	return res
}