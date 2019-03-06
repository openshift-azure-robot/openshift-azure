package validate

import (
	"errors"
	"reflect"
	"testing"

	pluginapi "github.com/openshift/openshift-azure/pkg/api/plugin/api"
)

func TestPluginTemplateValidate(t *testing.T) {
	expectedErrs :=
		[]error{
			errors.New(`invalid clusterVersion ""`),
			errors.New(`invalid imageOffer ""`),
			errors.New(`invalid imagePublisher ""`),
			errors.New(`invalid imageSKU ""`),
			errors.New(`invalid imageVersion ""`),
			errors.New(`invalid certificates.genevaLogging.key`),
			errors.New(`invalid certificates.genevaLogging.cert`),
			errors.New(`invalid certificates.genevaMetrics.key`),
			errors.New(`invalid certificates.genevaMetrics.cert`),
			errors.New(`invalid images.format ""`),
			errors.New(`invalid images.clusterMonitoringOperator ""`),
			errors.New(`invalid images.azureControllers ""`),
			errors.New(`invalid images.prometheusOperatorBase ""`),
			errors.New(`invalid images.prometheusBase ""`),
			errors.New(`invalid images.prometheusConfigReloaderBase ""`),
			errors.New(`invalid images.configReloaderBase ""`),
			errors.New(`invalid images.alertManagerBase ""`),
			errors.New(`invalid images.nodeExporterBase ""`),
			errors.New(`invalid images.grafanaBase ""`),
			errors.New(`invalid images.kubeStateMetricsBase ""`),
			errors.New(`invalid images.kubeRbacProxyBase ""`),
			errors.New(`invalid images.oAuthProxyBase ""`),
			errors.New(`invalid images.masterEtcd ""`),
			errors.New(`invalid images.controlPlane ""`),
			errors.New(`invalid images.node ""`),
			errors.New(`invalid images.serviceCatalog ""`),
			errors.New(`invalid images.sync ""`),
			errors.New(`invalid images.startup ""`),
			errors.New(`invalid images.templateServiceBroker ""`),
			errors.New(`invalid images.registry ""`),
			errors.New(`invalid images.router ""`),
			errors.New(`invalid images.registryConsole ""`),
			errors.New(`invalid images.ansibleServiceBroker ""`),
			errors.New(`invalid images.webConsole ""`),
			errors.New(`invalid images.console ""`),
			errors.New(`invalid images.etcdBackup ""`),
			errors.New(`invalid images.genevaImagePullSecret ""`),
			errors.New(`invalid images.genevaLogging ""`),
			errors.New(`invalid images.genevaTDAgent ""`),
			errors.New(`invalid images.genevaStatsd ""`),
			errors.New(`invalid images.metricsBridge ""`),
			errors.New(`invalid images.imagePullSecret ""`),
			errors.New(`invalid genevaLoggingSector ""`),
			errors.New(`invalid genevaLoggingAccount ""`),
			errors.New(`invalid genevaLoggingNamespace ""`),
			errors.New(`invalid genevaLoggingControlPlaneAccount ""`),
			errors.New(`invalid genevaLoggingControlPlaneEnvironment ""`),
			errors.New(`invalid genevaLoggingControlPlaneRegion ""`),
			errors.New(`invalid genevaMetricsAccount ""`),
			errors.New(`invalid genevaMetricsEndpoint ""`),
		}

	template := pluginapi.Config{}
	v := PluginAPIValidator{}
	errs := v.Validate(&template)
	if !reflect.DeepEqual(errs, expectedErrs) {
		t.Errorf("expected errors:")
		for _, err := range expectedErrs {
			t.Errorf("\t%v", err)
		}
		t.Error("received errors:")
		for _, err := range errs {
			t.Errorf("\t%v", err)
		}
	}
}
