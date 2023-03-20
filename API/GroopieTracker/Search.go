package GroopieTracker

import (
	"fmt"
	"main/Front/MoreContainer"
	"main/Front/MoreContainer/MemberContainer"
	"main/res"
	_var "main/res/var"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TODO Faire les filtres avec des if dans Search_Api()
// TODO mettre en Maj pour les cas "SOJA"

func SearchApi() {
	//fmt.Println(_var.Inputsaisie.Text)
	FilterApi()
	sort.Strings(_var.Resultat)
}

func SearchBand(UserInputStr string) []string {
	// Artists/band name
	var Bandlist []string

	for i := 0; i < len(_var.AData); i++ {
		if UserInputStr == _var.AData[i].Name {
			Bandlist = append(Bandlist, _var.AData[i].Name+res.L.SearchCategoryArtiste)
		} else if strings.Contains(_var.AData[i].Name, UserInputStr) == true {
			Bandlist = append(Bandlist, _var.AData[i].Name+res.L.SearchCategoryArtiste)
		} else if strings.EqualFold(_var.AData[i].Name, UserInputStr) == true {
			Bandlist = append(Bandlist, _var.AData[i].Name+res.L.SearchCategoryArtiste)
		} else if strings.Contains(_var.AData[i].Name, strings.ToTitle(UserInputStr)) == true {
			Bandlist = append(Bandlist, _var.AData[i].Name+res.L.SearchCategoryArtiste)
		}
	}
	return Bandlist
}

func SearchMember(UserInputStr string) []string {
	// members
	var Memberlist []string

	for x := range _var.AData {
		for y := range _var.AData[x].Members {
			if _var.AData[x].Members[y] == UserInputStr {
				Memberlist = append(Memberlist, _var.AData[x].Members[y]+res.L.SearchCategoryMember)
			} else if strings.Contains(_var.AData[x].Members[y], UserInputStr) == true {
				Memberlist = append(Memberlist, _var.AData[x].Members[y]+res.L.SearchCategoryMember)
			} else if strings.EqualFold(_var.AData[x].Members[y], UserInputStr) == true {
				Memberlist = append(Memberlist, _var.AData[x].Members[y]+res.L.SearchCategoryMember)
			} else if strings.Contains(_var.AData[x].Members[y], strings.ToTitle(UserInputStr)) == true {
				Memberlist = append(Memberlist, _var.AData[x].Members[y]+res.L.SearchCategoryMember)
			}
		}
	}
	return Memberlist
}

func SearchLocation(UserInputStr string) []string {
	// locations
	var Locationlist []string

	for _, rel := range _var.RData.Index {
		for location := range rel.DateLocation {
			_var.SelectLoclist = append(_var.SelectLoclist, location)
			if location == UserInputStr {
				Locationlist = append(Locationlist, location+res.L.SearchCategoryLocation)
			} else if strings.Contains(location, UserInputStr) == true {
				Locationlist = append(Locationlist, location+res.L.SearchCategoryLocation)
			} else if strings.EqualFold(location, UserInputStr) == true {
				Locationlist = append(Locationlist, location+res.L.SearchCategoryLocation)
			} else if strings.Contains(location, strings.ToTitle(UserInputStr)) == true {
				Locationlist = append(Locationlist, location+res.L.SearchCategoryLocation)
			}
		}
	}
	return Locationlist
}

func SearchConcDate(UserInputStr string) []string {
	// concert date
	var LocDatelist []string

	for _, rel := range _var.RData.Index {
		for _, datelist := range rel.DateLocation {
			for i := range datelist {
				if _var.MinMaxBool == true {
					LocDateTime, _ := time.Parse(_var.Layout, datelist[i])
					UserInputTime, _ := time.Parse(_var.Layout, UserInputStr)
					if LocDateTime.After(_var.UserMinDate) && LocDateTime.Before(_var.UserMaxDate) {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					} else if LocDateTime.Equal(UserInputTime) {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					}
				} else if _var.MinMaxBool == false {
					_var.MinMaxTimelist = append(_var.MinMaxTimelist, datelist[i])
					if datelist[i] == UserInputStr {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					} else if strings.Contains(datelist[i], UserInputStr) == true {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					} else if strings.EqualFold(datelist[i], UserInputStr) == true {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					} else if strings.Contains(datelist[i], strings.ToTitle(UserInputStr)) == true {
						LocDatelist = append(LocDatelist, datelist[i]+res.L.SearchCategoryConcertDate)
					}
				}
			}
		}
	}
	return LocDatelist
}

func SearchCreaDate(UserInputStr string) []string {
	// creation date
	var CreaDatelist []string

	for i := range _var.AData {
		if _var.MinMaxBool == true {
			creationDateInt := _var.AData[i].CreationDate
			creationDateStr := strconv.Itoa(_var.AData[i].CreationDate)
			_var.UserInputInt, _ = strconv.Atoi(UserInputStr)
			if creationDateInt >= _var.UserMin && creationDateInt <= _var.UserMax {
				if creationDateInt == _var.UserInputInt {
					CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
				} else {
					CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
				}
			}
		} else if _var.MinMaxBool == false {
			creationDateStr := strconv.Itoa(_var.AData[i].CreationDate)
			_var.MinMaxDatelist = append(_var.MinMaxDatelist, creationDateStr)
			if creationDateStr == UserInputStr {
				CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
			} else if strings.Contains(creationDateStr, UserInputStr) == true {
				CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
			} else if strings.EqualFold(creationDateStr, UserInputStr) == true {
				CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
			} else if strings.Contains(creationDateStr, strings.ToTitle(UserInputStr)) == true {
				CreaDatelist = append(CreaDatelist, creationDateStr+res.L.SearchCategoryCreationDate)
			}
		}
	}
	sort.Strings(_var.MinMaxDatelist)
	return CreaDatelist
}

func SearchAlbumDate(UserInputStr string) []string {
	// first album date
	var AlbDatelist []string

	for x := range _var.AData {
		if _var.MinMaxBool == true {
			AlbDateTime, _ := time.Parse(_var.Layout, _var.AData[x].FirstAlbum)
			if AlbDateTime.After(_var.UserMinDate) && AlbDateTime.Before(_var.UserMaxDate) {
				AlbDatelist = append(AlbDatelist, _var.AData[x].FirstAlbum+res.L.SearchCategoryAlbumDate)
			}
		} else if _var.MinMaxBool == false {
			if _var.AData[x].FirstAlbum == UserInputStr {
				AlbDatelist = append(AlbDatelist, _var.AData[x].FirstAlbum+res.L.SearchCategoryAlbumDate)
			} else if strings.Contains(_var.AData[x].FirstAlbum, UserInputStr) == true {
				AlbDatelist = append(AlbDatelist, _var.AData[x].FirstAlbum+res.L.SearchCategoryAlbumDate)
			} else if strings.EqualFold(_var.AData[x].FirstAlbum, UserInputStr) == true {
				AlbDatelist = append(AlbDatelist, _var.AData[x].FirstAlbum+res.L.SearchCategoryAlbumDate)
			} else if strings.Contains(_var.AData[x].FirstAlbum, strings.ToTitle(UserInputStr)) == true {
				AlbDatelist = append(AlbDatelist, _var.AData[x].FirstAlbum+res.L.SearchCategoryAlbumDate)
			}
		}
	}
	return AlbDatelist
}

func SelectFunc(s string) {
	fmt.Println()

	if s[len(s)-len(res.L.SearchCategoryArtiste):] == res.L.SearchCategoryArtiste {
		artist := s[:len(s)-len(res.L.SearchCategoryArtiste)]

		//search for artist/band place
		for i, a := range _var.AData {
			if a.Name == artist {

				fmt.Println(i)
				//open artist/band window
				MoreContainer.MoreWindow(i)
			}
		}

	} else if s[len(s)-len(res.L.SearchCategoryMember):] == res.L.SearchCategoryMember {
		member := s[:len(s)-len(res.L.SearchCategoryMember)]

		fmt.Println(member)
		//search for member place
		for aI, a := range _var.AData {
			for i, m := range a.Members {
				if member == m {

					//open the wikipedia of member
					MemberContainer.OpenWiki(aI, i)
				}
			}
		}
	} else if s[len(s)-len(res.L.SearchCategoryLocation):] == res.L.SearchCategoryLocation {
		fmt.Println("locations")
	} else if s[len(s)-len(res.L.SearchCategoryConcertDate):] == res.L.SearchCategoryConcertDate {
		fmt.Println("dates")
	} else if s[len(s)-len(res.L.SearchCategoryAlbumDate):] == res.L.SearchCategoryAlbumDate {
		fmt.Println("concerts")
	} else if s[len(s)-len(res.L.SearchCategoryCreationDate):] == res.L.SearchCategoryCreationDate {
		fmt.Println("creation")
	} else {
		fmt.Println("error")
	}
}
