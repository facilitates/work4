package serializer

type Member struct {
	Title 		string 	`json:"title"`
	Rank		int 	`json:"rank"`
	Views		int 	`json:"views"`
	VideoURL  	string	`json:"videourl"`
}

func BuildRankList(items []Member) (ranks []Member){
	for _, item := range items {
		rank := BuildMember(item)
		ranks = append(ranks, rank)
	}
	return ranks
}

func BuildMember(item Member) Member{
	return Member {
		Title: item.Title,
		Rank: item.Rank,
		Views: item.Views,
		VideoURL: item.VideoURL,
	}
}