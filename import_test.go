package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImportService_ImportRepositoryFromGitHub(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/import/github", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"id": 27,
				"name": "my-repo",
				"full_path": "/root/my-repo",
				"full_name": "Administrator / my-repo",
				"refs_url": "/root/my-repo/refs",
				"import_source": "my-github/repo",
				"import_status": "scheduled",
				"human_import_status_name": "scheduled",
				"provider_link": "/my-github/repo",
				"relation_type": null,
				"import_warning": null
			}
		`)
	})

	want := &GitHubImport{
		ID:                    27,
		Name:                  "my-repo",
		FullPath:              "/root/my-repo",
		FullName:              "Administrator / my-repo",
		RefsUrl:               "/root/my-repo/refs",
		ImportSource:          "my-github/repo",
		ImportStatus:          "scheduled",
		HumanImportStatusName: "scheduled",
		ProviderLink:          "/my-github/repo",
	}

	opt := &ImportRepositoryFromGitHubOptions{
		PersonalAccessToken: Ptr("token"),
		RepoID:              Ptr(34),
		TargetNamespace:     Ptr("root"),
	}

	gi, resp, err := client.Import.ImportRepositoryFromGitHub(opt)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gi)

	gi, resp, err = client.Import.ImportRepositoryFromGitHub(opt, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gi)
}

func TestImportService_CancelGitHubProjectImport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/import/github/cancel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"id": 27,
				"name": "my-repo",
				"full_path": "/root/my-repo",
				"full_name": "Administrator / my-repo",
				"import_source": "my-github/repo",
				"import_status": "scheduled",
				"human_import_status_name": "scheduled",
				"provider_link": "/my-github/repo"
			}
		`)
	})

	want := &CancelledGitHubImport{
		ID:                    27,
		Name:                  "my-repo",
		FullPath:              "/root/my-repo",
		FullName:              "Administrator / my-repo",
		ImportSource:          "my-github/repo",
		ImportStatus:          "scheduled",
		HumanImportStatusName: "scheduled",
		ProviderLink:          "/my-github/repo",
	}

	opt := &CancelGitHubProjectImportOptions{
		ProjectID: Ptr(27),
	}

	cgi, resp, err := client.Import.CancelGitHubProjectImport(opt)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, cgi)

	cgi, resp, err = client.Import.CancelGitHubProjectImport(opt, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, cgi)
}

func TestImportService_ImportGitHubGistsIntoGitLabSnippets(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/import/github/gists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	opt := &ImportGitHubGistsIntoGitLabSnippetsOptions{PersonalAccessToken: Ptr("token")}

	resp, err := client.Import.ImportGitHubGistsIntoGitLabSnippets(opt)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Import.ImportGitHubGistsIntoGitLabSnippets(opt, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
}
