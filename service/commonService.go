package service

import (
	"fmt"
	"test/models"
	"test/repository"
)

type CommonService struct {
	CommonRepo repository.CommonRepo
}

func NewCommonService(commonRepo repository.CommonRepo) CommonService {
	return CommonService{CommonRepo: commonRepo}
}

func (cs *CommonService) Task1Service(body models.Task1Model) ([][]int, error) {
	arr := body.Numbers

	if len(arr) == 0 {
		return nil, fmt.Errorf("the array is empty")
	}
	result := [][]int{}
	duplicate := make(map[int]int)

	for i:=0; i<len(arr); i++ {
		complement := body.Target - arr[i]
		if idx, found := duplicate[complement]; found {
			result = append(result, []int{idx, i})
		}
		duplicate[arr[i]] = i
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no pairs found matching the target sum")
	}
	fmt.Println("The response is : ", result)
	return result, nil
}
