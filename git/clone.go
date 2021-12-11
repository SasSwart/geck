package git

import git "github.com/libgit2/git2go/v33"

func Clone(repo, path, branch string, sshCredential SSHCredential) (*git.Repository, error) {
	callbacks := git.RemoteCallbacks{
		CertificateCheckCallback: CertificateCheckCallback,
		CredentialsCallback:      sshCredential.GetCallback(),
	}
	return git.Clone(repo, path, &git.CloneOptions{
		CheckoutBranch: branch,
		CheckoutOptions: git.CheckoutOptions{
			Strategy: git.CheckoutForce,
		},
		FetchOptions: git.FetchOptions{
			RemoteCallbacks: callbacks,
		},
	})
}
