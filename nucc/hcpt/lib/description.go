package lib

import (
  "os"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "nucc.hcpt",
      Version: "2015-01",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  return []string{ "Code", "Type", "Classification", "Specialization", "Definition", "Notes" }, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://www.nucc.org/images/stories/CSV/nucc_taxonomy_150.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}