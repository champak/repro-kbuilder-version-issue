package v1

import (
	"conv/api/v1alpha1"

	debugLog "log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var (
	convLog = ctrl.Log.WithName("conv")
)

func (src *Repro) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha1.Repro)

	debugLog.Printf("ConvertTo working on %s \n", src.ObjectMeta.Name)
	debugLog.Printf("ConvertTo working on Pod %s \n", src.Spec.Graph.Pod.ObjectMeta.Name)

	debugLog.Printf("ConvertTo working on Pod %s attr %s \n", src.Spec.Graph.Pod.ObjectMeta.Name, src.Spec.Foo)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.Foo = src.Spec.Foo

	return nil
}

func (dst *Repro) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha1.Repro)

	debugLog.Printf("ConvertFrom : working on %s \n", src.ObjectMeta.Name)

	debugLog.Printf("ConvertFrom working on Pod %s attr %s \n", src.Spec.Graph.Pod.ObjectMeta.Name, src.Spec.Foo)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.Foo = src.Spec.Foo

	return nil

}
