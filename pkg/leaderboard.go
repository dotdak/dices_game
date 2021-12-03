package pkg

type User struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type LeaderBoard struct {
	Data []*User
	Size int
}

func (b *LeaderBoard) Push(u *User) {
	n := len(b.Data)
	if n == 0 {
		b.Data = append(b.Data, u)
		return
	}

	if n == b.Size && u.Score < (b.Data)[n-1].Score {
		return
	}
	b.Data = append(b.Data, u)
	for k, v := range (b.Data)[:n] {
		if v.Score < u.Score {
			copy((b.Data)[k+1:n+1], (b.Data)[k:n])
			(b.Data)[k] = u
			break
		}
	}
	if len(b.Data) > b.Size {
		(b.Data)[b.Size] = nil
		b.Data = (b.Data)[:b.Size]
	}
}
