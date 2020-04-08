package models

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type HomeBlockParam struct {
	//Id int
	//Title string
	//Tags []TagLink
	//Short string
	//Content string
	//Author string
	//CreateTime string
	Article       *Article
	TagLinks      []*TagLink
	CreateTimeStr string
	//查看文章的地址
	Link string
	//修改文章的地址
	UpdateLink string
	DeleteLink string
	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

func createTagsLinks(tags string) []*TagLink {
	//var tagLinks = make([]*TagLink, 0, strings.Count(tags, "&"))
	var tagLinks []*TagLink
	tagsList := strings.Split(tags, "&")
	for _, tag := range tagsList {
		tagLinks = append(tagLinks, &TagLink{tag, "/?tag=" + tag})
	}
	return tagLinks
}

func GenHomeBlocks(articleList []*Article, isLogin bool) (ret []*HomeBlockParam) {
	ret = make([]*HomeBlockParam, 0, len(articleList))

	for _, art := range articleList {
		homeParam := HomeBlockParam{
			Article: art,
			IsLogin: isLogin,
		}
		homeParam.TagLinks = createTagsLinks(art.Tags)
		homeParam.CreateTimeStr = art.CreateTime
		homeParam.Link = fmt.Sprintf("/show/%d", art.Id)
		homeParam.UpdateLink = fmt.Sprintf("/article/update?id=%d", art.Id)
		homeParam.DeleteLink = fmt.Sprintf("/article/delete?id=%d", art.Id)

		ret = append(ret, &homeParam)
	}
	return ret
}

func MakeHomeBlocks(articles []*Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		homeParam := HomeBlockParam{
			Article: art,
			IsLogin: isLogin,
		}
		homeParam.TagLinks = createTagsLinks(art.Tags)
		homeParam.CreateTimeStr = art.CreateTime
		homeParam.Link = fmt.Sprintf("/show/%d", art.Id)
		homeParam.UpdateLink = fmt.Sprintf("/article/update?id=%d", art.Id)
		homeParam.DeleteLink = fmt.Sprintf("/article/delete?id=%d", art.Id)
		//homeParam := HomeBlockParam{}
		//homeParam.Id	= art.Id
		//homeParam.Title = art.Title
		//homeParam.Tags = createTagsLinks(art.Tags)
		//fmt.Println("tag-->", art.Tags)
		//homeParam.Short = art.Short
		//homeParam.Content = art.Content
		//homeParam.Author = art.Author
		//homeParam.CreateTime = utils.SwitchTimeStampToData(art.CreateTime)
		//homeParam.Link = "/show/" + strconv.Itoa(art.Id)
		//homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		//homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		//homeParam.IsLogin = isLogin

		t, _ := template.ParseFiles("templates/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}
