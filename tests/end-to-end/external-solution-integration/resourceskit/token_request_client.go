package resourceskit

import (
	"github.com/kyma-project/kyma/components/connection-token-handler/pkg/apis/applicationconnector/v1alpha1"
	"github.com/kyma-project/kyma/components/connection-token-handler/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"time"
)

type TokenRequestClient interface {
	CreateTokenRequest(appName string) (*v1alpha1.TokenRequest, error)
	GetTokenRequest(appName string, options v1.GetOptions) (*v1alpha1.TokenRequest, error)
	DeleteTokenRequest(appName string, options *v1.DeleteOptions) error
}

type tokenRequestClient struct {
	client    *versioned.Clientset
	namespace string
}

func NewTokenRequestClient(config *rest.Config, namespace string) (TokenRequestClient, error) {
	clientSet, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &tokenRequestClient{client: clientSet, namespace: namespace}, nil
}

func (t *tokenRequestClient) CreateTokenRequest(appName string) (*v1alpha1.TokenRequest, error) {
	tokenRequest := &v1alpha1.TokenRequest{
		TypeMeta:   v1.TypeMeta{Kind: "Application", APIVersion: v1alpha1.SchemeGroupVersion.String()},
		ObjectMeta: v1.ObjectMeta{Name: appName, Namespace: t.namespace},
		Status: v1alpha1.TokenRequestStatus{
			ExpireAfter: v1.NewTime(time.Now().Add(5 * time.Minute)),
		},
	}

	return t.client.ApplicationconnectorV1alpha1().TokenRequests(t.namespace).Create(tokenRequest)
}

func (t *tokenRequestClient) GetTokenRequest(appName string, options v1.GetOptions) (*v1alpha1.TokenRequest, error) {
	return t.client.ApplicationconnectorV1alpha1().TokenRequests(t.namespace).Get(appName, options)
}

func (t *tokenRequestClient) DeleteTokenRequest(appName string, options *v1.DeleteOptions) error {
	return t.client.ApplicationconnectorV1alpha1().TokenRequests(t.namespace).Delete(appName, options)
}
