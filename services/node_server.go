package services

import "nable.gin/app/models"

//迭代实现
func Iteration(allCate []models.Node, pid int) []models.Node {
	task := []int{}
	task = append(task, 0)
	res := []models.Node{}
	hasChild := false
	var parent int = pid
	for len(task) > 0 {
		hasChild = false
		for k, v := range allCate {
			if v.ID == -1 {
				continue
			}
			if parent == v.Pid {
				res = append(res, v)
				task = append(task, v.ID)
				allCate[k].ID = -1 //奖该数据删除
				hasChild = true
				parent = v.ID
				break
			}
		}

		if !hasChild {
			end := len(task) - 1
			task = task[0:end] //将该数据删除
			//继续找它上级得其他子类
			if len(task) > 0 {
				end = end - 1
				parent = task[end]
			}

		}

	}
	return res
}