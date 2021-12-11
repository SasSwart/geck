package git

import git "github.com/libgit2/git2go/v33"

func CertificateCheckCallback(cert *git.Certificate, valid bool, hostname string) error {
	return nil //git.CertificateCheckCallback(cert, valid, hostname)
}

type SSHCredential struct {
	PrivateKeyPath string
	PublicKeyPath  string
	PassPhrase     string
}

func (s SSHCredential) GetCallback() func(string, string, git.CredentialType) (*git.Credential, error) {
	return func(url string, username string, allowedTypes git.CredentialType) (*git.Credential, error) {
		return git.NewCredentialSSHKey(username, s.PublicKeyPath, s.PrivateKeyPath, s.PassPhrase)
		//git.NewCredentialSSHKeyFromAgent(username)
	}
}
