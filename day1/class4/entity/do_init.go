package entity

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicMap map[int64]*Topic
	postMap  map[int64][]*Post
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

// 初始化话题map
func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "/topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.ID] = &topic
	}
	topicMap = topicTmpMap
	return nil
}

// 初始化评论map
func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "/post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts, ok := postTmpMap[post.ParentID] // 检查当前评论是否已经存在
		if !ok {
			postTmpMap[post.ParentID] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentID] = posts
	}
	postMap = postTmpMap
	return nil
}
