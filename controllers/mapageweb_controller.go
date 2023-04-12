/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	//v1 "kubebuilder-project/api/v1"
	v1 "kubebuilder-operator-test/projet/api/v1"
)

// MaPageWebReconciler reconciles a MaPageWeb object
type MaPageWebReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test.kubebuilder-operator-test,resources=mapagewebs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test.kubebuilder-operator-test,resources=mapagewebs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test.kubebuilder-operator-test,resources=mapagewebs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
//+kubebuilder:object:root=true

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MaPageWeb object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *MaPageWebReconciler) Reconcile(ctx context.Context, req ctrl.Request, MaPageWeb *v1.MaPageWeb) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	// SetupWithManager sets up the controller with the Manager.

	//https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta
	//https://pkg.go.dev/k8s.io/api/core/v1#ConfigMap

	//Data to Store in the ConfigMap (a web content)
	data := map[string]string{
		"index.html": "<html><body><h1> " + MaPageWeb.Spec.Contenu + " </h1></body></html>",
	}

	//Creating ConfigMap
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: MaPageWeb.Name + "-config",
		},
		Data: data,
	}

	// Set MyResource instance as the owner and controller of the ConfigMap
	if err := ctrl.SetControllerReference(configMap, r.Scheme); err != nil {
		return err
	}
	// Create or update the ConfigMap
	err := r.Create(ctx, configMap)
	if err != nil && !errors.IsAlreadyExists(err) {

		return err
	}
	var Preference = MaPageWeb.Spec.Pref
	var Repliques int32

	switch Preference {

	case "Critique":

		Repliques = 6

	case "Important":
		Repliques = 4

	case "Normal":
		Repliques = 2

	}

	Deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: MaPageWeb.Name,
		},
		Spec: appsv1.DeploymentSpec{

			Replicas: Repliques,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": MaPageWeb.Application,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": MaPageWeb.Application,
					},
				},
				Spec: v1.PodSpec{
					containers: v1.Containers{
						Name: MaPageWeb.Spec.Application,

						Image: MaPageWeb.Spec.Application,

						Ports: v1.ContainerPort{
							ContainerPort: 80,
						},
						VolumeMounts: v1.VolumeMount{

							Name:      MaPageWeb.Name + "-storage",
							MountPath: "/usr/share/" + MaPageWeb.Application + "/html",
						},
					},
					volumes: v1.Volume{

						Name: MaPageWeb.Name + "-storage",

						ConfigMap: v1.ConfigMapVolumeSource{
							Name: MaPageWeb.Application + "-config",

							Items: v1.KeyToPath{
								Key:  "index.html",
								Path: "index.html",
							},
						},
					},
				},
			},
		},
	}

	return ctrl.Result{}, nil
}
func (r *MaPageWebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testv1.MaPageWeb{}).
		Complete(r)
}
