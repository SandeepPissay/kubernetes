/*
Copyright 2019 The Kubernetes Authors.

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

package plugins

import (
	"fmt"
	"k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CnsCsiDriverName = "csi.vsphere.vmware.com"
	VSphereVolumeInTreePluginName = "kubernetes.io/vsphere-volume"
)

var _ InTreePlugin = &vSphereVolumeCSITranslator{}

type vSphereVolumeCSITranslator struct{}

func (vSphereVolumeCSITranslator) TranslateInTreeStorageClassToCSI(sc *storage.StorageClass) (*storage.StorageClass, error) {
	fmt.Printf("TranslateInTreeStorageClassToCSI: implement me")
	return sc, nil
}

func (vSphereVolumeCSITranslator) TranslateInTreeInlineVolumeToCSI(volume *v1.Volume) (*v1.PersistentVolume, error) {
	fmt.Printf("TranslateInTreeInlineVolumeToCSI: implement me")
	if volume == nil || volume.AzureFile == nil {
		return nil, fmt.Errorf("volume is nil or AWS EBS not defined on volume")
	}

	vsphereSource := volume.VsphereVolume

	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: volume.Name,
		},
		Spec: v1.PersistentVolumeSpec{
			PersistentVolumeSource: v1.PersistentVolumeSource{
				CSI: &v1.CSIPersistentVolumeSource{
					VolumeHandle:     vsphereSource.VolumePath,
					VolumeAttributes: map[string]string{},
				},
			},
			AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
		},
	}
	return pv, nil
}

func (vSphereVolumeCSITranslator) TranslateInTreePVToCSI(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	fmt.Printf("TranslateInTreePVToCSI: implement me")
	return pv, nil
}

func (vSphereVolumeCSITranslator) TranslateCSIPVToInTree(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	fmt.Printf("TranslateCSIPVToInTree: implement me")
	return pv, nil
}

func (vSphereVolumeCSITranslator) CanSupport(pv *v1.PersistentVolume) bool {
	fmt.Printf("CanSupport: implement me")
	return pv != nil && pv.Spec.VsphereVolume != nil
}

func (vSphereVolumeCSITranslator) CanSupportInline(vol *v1.Volume) bool {
	fmt.Print("CanSupportInline: implement me")
	return vol != nil && vol.VsphereVolume != nil
}

func (vSphereVolumeCSITranslator) GetInTreePluginName() string {
	return VSphereVolumeInTreePluginName
}

func (vSphereVolumeCSITranslator) GetCSIPluginName() string {
	return CnsCsiDriverName
}

func (vSphereVolumeCSITranslator) RepairVolumeHandle(volumeHandle, nodeID string) (string, error) {
	return volumeHandle, nil
}

// NewVSphereVolumeCSITranslator returns a new instance of vSphereVolumeCSITranslator
func NewVSphereVolumeCSITranslator() InTreePlugin {
	return &vSphereVolumeCSITranslator{}
}


