package git

import git "github.com/libgit2/git2go/v33"

func credentialsCallback(url string, username string, allowedTypes git.CredentialType) (*git.Credential, error) {
	return git.NewCredentialSSHKey(username, "/home/sas/.ssh/id_rsa.pub", "/home/sas/.ssh/id_rsa", "")
	//git.NewCredentialSSHKeyFromAgent(username)
}

func CertificateCheckCallback(cert *git.Certificate, valid bool, hostname string) error {
	return nil //git.CertificateCheckCallback(cert, valid, hostname)
}
