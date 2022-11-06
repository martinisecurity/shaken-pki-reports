package repository_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/repository"
)

func TestListBucketResultFromXML(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    *repository.ListBucketResult
		wantErr bool
	}{
		{ // correct ListBucketResult
			name: "correct ListBucketResult",
			args: args{
				data: `<?xml version='1.0' encoding='UTF-8'?>
<ListBucketResult xmlns='http://doc.s3.amazonaws.com/2006-03-01'>
		<Name>martini-prod-certs</Name>
		<Prefix></Prefix>
		<Marker></Marker>
		<IsTruncated>false</IsTruncated>
		<Contents>
				<Key>2884/BiAXt2iK2Pw-.pem</Key>
				<Generation>1667417869151511</Generation>
				<MetaGeneration>2</MetaGeneration>
				<LastModified>2022-11-02T19:37:49.185Z</LastModified>
				<ETag>"a5996b3fcfbfa9f323ab8c27141d4e4c"</ETag>
				<Size>2253</Size>
		</Contents>
		<Contents>
				<Key>2884/CR0M-J_8cVtU.pem</Key>
				<Generation>1666731452787202</Generation>
				<MetaGeneration>2</MetaGeneration>
				<LastModified>2022-10-25T20:57:32.807Z</LastModified>
				<ETag>"56bbbe637fcd5b2230ff8cbb3b17d887"</ETag>
				<Size>2273</Size>
		</Contents>
		<Contents>
				<Key>2884/DmEER_UmW8Vq.pem</Key>
				<Generation>1667415687925282</Generation>
				<MetaGeneration>2</MetaGeneration>
				<LastModified>2022-11-02T19:01:27.956Z</LastModified>
				<ETag>"4fa92a64130291a3d140b8fa227fc29d"</ETag>
				<Size>2294</Size>
		</Contents>
</ListBucketResult>`,
			},
			want: &repository.ListBucketResult{
				Contents: []repository.ContentResult{
					{
						Key: "2884/BiAXt2iK2Pw-.pem",
					},
					{
						Key: "2884/CR0M-J_8cVtU.pem",
					},
					{
						Key: "2884/DmEER_UmW8Vq.pem",
					},
				},
			},
			wantErr: false,
		},
		{ // incorrect ListBucketResult
			name: "incorrect ListBucketResult",
			args: args{
				data: `some text`,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.data)
			got, err := repository.ListBucketResultFromXML(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBucketResultFromXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucketResultFromXML() = %v, want %v", got, tt.want)
			}
		})
	}
}
