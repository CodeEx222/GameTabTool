package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexDefine_ContainTag(t *testing.T) {
	// 测试包含标签的情况
	index := IndexDefine{
		Group: []string{"tag1", "tag2", "tag3"},
	}
	assert.True(t, index.ContainTag("tag2"), "ContainTag failed for existing tag")

	// 测试不包含标签的情况
	assert.False(t, index.ContainTag("tag4"), "ContainTag failed for non-existing tag")

	// 测试空标签组的情况
	index = IndexDefine{
		Group: []string{},
	}
	assert.False(t, index.ContainTag("tag1"), "ContainTag failed for empty group")

	// 测试空字符串标签的情况
	index = IndexDefine{
		Group: []string{"", "tag2"},
	}
	assert.True(t, index.ContainTag(""), "ContainTag failed for empty string tag")

}
