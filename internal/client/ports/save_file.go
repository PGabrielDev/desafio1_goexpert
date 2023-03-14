package ports

type ISaveFile interface {
	SaveFile(dolar string) error
}
