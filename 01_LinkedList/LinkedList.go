package main

import (
	"fmt"
)


// Chapter singly Linked List
type Chapter struct{
	Order int
	Title string
	Text *Content
	Next *Chapter
}

// Content singly Linked List
type Content struct{
	Page int
	Title string
	Text string
	Next *Content
}


// WriteNovel Singly　Linked List
func WriteNovel()*Chapter{
	fmt.Println("Rain")
	// initHead
	head := &Chapter{}

	chapter01 := &Chapter{
		Order: 1,
		Title: "青铜巨棺",
	}
	chapter02 := &Chapter{
		Order: 2,
		Title: "素问",
	}
	chapter03 := &Chapter{
		Order: 3,
		Title: "今昔",
	}
	chapter04 := &Chapter{
		Order: 4,
		Title:"荒古铜刻",
	}

	content01 := []string{"生命是世界最伟大的奇迹。","四方上下宇宙。","宇虽有实，　而无定处可求。","往古来今曰宙。"}
	content02 := []string{"上古之人，春秋皆满百岁，而动作不衰。","关于上古时期， 并没有详细而准确地记载，","对于今人来说那是一段充满无尽迷雾的古史， 让人遐想无限。","清风拂动， 院中几株梧桐在轻轻的摇曳， 繁茂的枝叶发出簌簌的声响， 清新的空气从窗外迎面吹来。"}
	content03 := []string{"叶凡虽然谈不上事业有成， 但因为一些原因和经历， 如今手里也有些资本，","前不久手里恰好买了一辆奔驰车。 就价格而言要高于刘云志的那款丰田车","但如果以此来评定身份的话，他感觉相当庸俗","十几分钟后，叶凡驱车来到聚会点--海上明月城"}
	content04 := []string{"夜幕悄悄降临， 道路两旁得霓虹灯闪烁，夜间的都市依然散发着无尽的活力，","一座座摩天大楼鳞次栉比， 耸入高空中。", "若从太空往下望去，这一切显得微不住道，不过方寸之地", "这是一个极不平静的夜晚， 注定难以平静，地面上数十个主监控室彻底锁定漆黑的苍穹。"}
	
	chapter01.Text = WriteContent(content01, chapter01.Title)
	chapter02.Text = WriteContent(content02, chapter02.Title)
	chapter03.Text = WriteContent(content03, chapter03.Title)
	chapter04.Text = WriteContent(content04, chapter04.Title)

	InsertChapter(head, chapter01)
	InsertChapter(head, chapter02)
	InsertChapter(head, chapter03)
	InsertChapter(head, chapter04)
	return head
}

// WriteContent Write chapter content
func WriteContent(contents []string, chapter string)*Content{
	head := &Content{}
	for i, c := range contents{
		content := &Content{
			Page: i,
			Text: c,
			Title: chapter,
		}
		InsertContent(head, content)
	}
	return head
}

// InsertChapter insert chapter in order
func InsertChapter(head *Chapter, chapter *Chapter){
	temp:=head
	for {
		if temp.Next == nil{
			break
		} else if temp.Next.Order == chapter.Order{
				return
		} else if temp.Next.Order > chapter.Order{
			break
		}
		temp = temp.Next
	}
	chapter.Next = temp.Next
	temp.Next = chapter
}

// InsertContent Insert Content in order
func InsertContent(head *Content, content *Content){
	temp := head
	for {
		if temp.Next == nil{
			break
		} else if temp.Next.Page == content.Page{
			return
		} else if temp.Next.Page > content.Page{
			break
		}
		temp = temp.Next
	}
	content.Next = temp.Next
	temp.Next = content
}

// ReadNovel start read novel
func ReadNovel(novel *Chapter){
	fmt.Println("<=== Start Read Novel ===>")
	chapter := novel.Next
	for {
		if chapter == nil{
			break
		}
		fmt.Printf("\n第%d章 %s\n",chapter.Order, chapter.Title)
		content := chapter.Text
		for {
			if content.Next == nil{
				break
			}
			if content.Text != ""{
				fmt.Println(content.Text)
			}
			content = content.Next
		}
		chapter = chapter.Next
	}

	fmt.Println("<========= END =========>")

}