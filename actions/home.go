package actions

import (
	"sort"

	"encoding/json"
	"github.com/gobuffalo/buffalo"
	"premutations_api/models"
)

func V1PremutationsInit(c buffalo.Context) error {
	data := &[]int{}
	if err := c.Bind(data); err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}
	sort.Ints(*data)
	dataForSave, err := json.Marshal(*data)
	if err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}
	prem := models.Premutation{Slice: string(dataForSave)}
	err = models.DB.Create(&prem)
	if err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}
	c.Session().Set("id", prem.ID)
	return c.Render(200, r.JSON(*data))
}

func V1PremutationsNext(c buffalo.Context) error {
	premIdTemp := c.Session().Get("id")
	if premIdTemp == nil {
		return c.Render(200, r.JSON(""))
	}
	premId, ok := premIdTemp.(int)

	if !ok {
		c.Logger().Error("type assertion error in V1PremutationsNext")
	}
	if premId == 0 {
		return c.Render(200, r.JSON([]int{}))
	}
	prem := models.Premutation{}
	err := models.DB.Find(&prem, premId)
	if err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}

	res := []int{}
	err = json.Unmarshal([]byte(prem.Slice), &res)
	if err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}
	haveRes := NextSet(res)
	dataForSave, err := json.Marshal(res)
	if err != nil {
		c.Logger().Error(err)
		return c.Render(500, r.JSON(""))
	}

	prem.Slice = string(dataForSave)
	err = models.DB.Save(&prem)
	if haveRes {
		return c.Render(200, r.JSON(res))
	} else {
		models.DB.Destroy(&prem)
		c.Session().Set("id", 0)
		return c.Render(200, r.JSON([]int{}))
	}
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func NextSet(a []int) bool {
	n := len(a)
	if n == 0 {
		return false
	}
	j := n - 2
	for ; j != -1 && a[j] >= a[j+1]; j-- {
	}
	if j == -1 {
		return false
	}
	k := n - 1
	for a[j] >= a[k] {
		k--
	}
	swap(a, j, k)
	l := j + 1
	r := n - 1
	for l < r {
		l++
		r--
		swap(a, l, r)
	}
	return true
}
