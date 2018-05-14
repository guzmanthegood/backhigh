package models

import (
    "time"
)

type Highlight struct {
    Id	        int
    Created_At  time.Time
    Title       string
    Country     string
    Price       float32
    Content     string
}

func AllHighlights() ([]*Highlight, error) {
    rows, err := db.Query("SELECT * FROM highlights")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    hls := make([]*Highlight, 0)
    for rows.Next() {
        hl := new(Highlight)
        err := rows.Scan(&hl.Id, &hl.Created_At, &hl.Title, &hl.Country, &hl.Price, &hl.Content)
        if err != nil {
            return nil, err
        }
        hls = append(hls, hl)
    }


    if err = rows.Err(); err != nil {
        return nil, err
    }

    return hls, nil
}