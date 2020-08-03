package controller

import (
	"reflect"
	"testing"
	"triproute/pkg/repository"
)

func TestGetBestRoute(t *testing.T) {
	type args struct {
		get RequestBestRoute
		r   *repository.Repository
	}
	tests := []struct {
		name    string
		args    args
		want    ResponseBestRoute
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBestRoute(tt.args.get, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBestRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBestRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
