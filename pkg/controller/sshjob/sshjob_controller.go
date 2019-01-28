package sshjob

import (
	"context"
  // "fmt"
	// shellv1alpha1 "ssh-operator/pkg/apis/shell/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_sshjob")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SSHJob Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSSHJob{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("sshjob-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SSHJob
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner SSHJob
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &corev1.Pod{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSSHJob{}

// ReconcileSSHJob reconciles a SSHJob object
type ReconcileSSHJob struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SSHJob object and makes changes based on the state read
// and what is in the SSHJob.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSSHJob) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling SSHJob")

	// Fetch the SSHJob instance
	instance := &corev1.Pod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}


  annotationIn := instance.ObjectMeta.Annotations["ssh.in"]
  annotationOut := instance.ObjectMeta.Annotations["ssh.out"]

	// Define a new Pod object
	pod := newPodForCR(instance)

  // Set SSHJob instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)

  if len(annotationIn) > 0 {
  	if err != nil && errors.IsNotFound(err) {
  		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
  		err = r.client.Create(context.TODO(), pod)
  		if err != nil {
  			return reconcile.Result{}, err
  		}
      // MIGHT need to put logic here or return an error - TEST

  		// Pod created successfully - don't requeue
  		return reconcile.Result{}, nil
  	} else if err != nil {
  		return reconcile.Result{}, err
  	}

    // Move this logic to container
    // if len(annotationOut) == 0 {
    //   fmt.Println("Here")
    //   annotationsWorker := found.ObjectMeta.Annotations["ssh"]
    //   fmt.Println(annotationsWorker)
    //   if len(annotationsWorker) > 0 {
    //     // copy the annotation
    //     m := make(map[string]string)
    //     currentAnnotations := instance.GetAnnotations()
    //     for k, v := range currentAnnotations {
    //       m[k] = v
    //     }
    //     m["ssh.out"] = annotationsWorker
    //
    //     instance.SetAnnotations(m)
    //
    //     err = r.client.Update(context.Background(), instance)
    //     if err != nil {
    //       return reconcile.Result{}, nil
    //     }
    //   	reqLogger.Info("Adding annotation", "Pod.Namespace", instance.Namespace, "Pod.Name", instance.Name)
    //   	return reconcile.Result{}, nil
    //    } else {
    //      // CREATE NEW Error here
    //   }
    // }

    // DELETE --v 2 lines
  	// Pod already exists - don't requeue
  	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
  	return reconcile.Result{}, nil
  } else {
    if len(annotationOut) > 0 {
      reqLogger.Info("Remove ssh annotation", "Pod.Namespace", instance.Namespace, "Pod.Name", instance.Name)
      m := make(map[string]string)
      currentAnnotations := instance.GetAnnotations()
      for k, v := range currentAnnotations {
        if k != "ssh.out" {
          m[k] = v
        }
      }
      instance.SetAnnotations(m)
      err = r.client.Update(context.Background(), instance)
      if err != nil {
        return reconcile.Result{}, nil
      }
    }
    // Dont forget to strip out labels
    if err != nil && errors.IsNotFound(err) {
      return reconcile.Result{}, nil
    } else if err != nil {
      return reconcile.Result{}, err
    }
    reqLogger.Info("Deleting a Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name, "Annotation", annotationIn)
    err = r.client.Delete(context.TODO(), pod)
    if err != nil {
      return reconcile.Result{}, err
    }
    return reconcile.Result{}, nil
  }
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *corev1.Pod) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tmate-" + cr.Name,//"test-" + cr.Name + "-pod",
			Namespace: "tmate", // cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
      ServiceAccountName: "ssh-operator",
      Volumes: []corev1.Volume{
        {
          Name: "config",
          VolumeSource: corev1.VolumeSource{
            ConfigMap: &corev1.ConfigMapVolumeSource{
              LocalObjectReference: corev1.LocalObjectReference{
                Name: "tmate-init",
              },
            },
          },
        },
        {
          Name: "podinfo",
          VolumeSource: corev1.VolumeSource{
            DownwardAPI: &corev1.DownwardAPIVolumeSource{
              Items: []corev1.DownwardAPIVolumeFile{
                {
                  Path: "name",
                  FieldRef: &corev1.ObjectFieldSelector{
                    FieldPath: "metadata.name",
                  },
                },
              },
            },
          },
        },
      },
			Containers: []corev1.Container{
				{
					Name:    "tmate-client",
					Image:   "cgetzen/tmate-kubectl:1.13.2",
					Command: []string{"sh", "/etc/tmate-init/init.sh"},
          VolumeMounts: []corev1.VolumeMount{
            {
              Name: "config",
              ReadOnly: true,
              MountPath: "/etc/tmate-init",
            },
            {
              Name: "podinfo",
              ReadOnly: false,
              MountPath: "/etc/podinfo",
            },
          },
          LivenessProbe: &corev1.Probe{
            Handler: corev1.Handler{
              Exec: &corev1.ExecAction{
                Command: []string{"tmate", "-S", "/tmp/tmate.sock", "wait", "tmate-ready"},
              },
            },
            InitialDelaySeconds: 10,
          },
          Env: []corev1.EnvVar{
            {
              Name: "DEST_POD",
              Value: cr.Name,
            },
            // {
            //   Name: "DEST_CONTAINER",
            //   Value: cr.Spec.Container,
            // },
            {
              Name: "DEST_NAMESPACE",
              Value: cr.Namespace,
            },
          },
				},
			},
		},
	}
}
