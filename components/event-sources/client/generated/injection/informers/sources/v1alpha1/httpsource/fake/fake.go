// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	"context"

	fake "github.com/kyma-project/kyma/components/event-sources/client/generated/injection/informers/factory/fake"
	httpsource "github.com/kyma-project/kyma/components/event-sources/client/generated/injection/informers/sources/v1alpha1/httpsource"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = httpsource.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Sources().V1alpha1().HTTPSources()
	return context.WithValue(ctx, httpsource.Key{}, inf), inf.Informer()
}
