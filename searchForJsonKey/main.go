package main

import (
	"encoding/json"
	"log"
	"strings"
)

type SysdigAlertsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Alerts []struct {
			Labels struct {
				Severity           string `json:"severity"`
				Alertname          string `json:"alertname"`
				KubeNamespaceName  string `json:"kube_namespace_name,omitempty"`
				KubeDeploymentName string `json:"kube_deployment_name,omitempty"`
				KubeClusterName    string `json:"kube_cluster_name,omitempty"`
				KubeNodeName       string `json:"kube_node_name,omitempty"`
				KubePodName        string `json:"kube_pod_name,omitempty"`
				Type               string `json:"kube_namespace_label_type,omitempty"`
			} `json:"labels,omitempty"`
			Annotations struct {
				Description string `json:"description"`
			} `json:"annotations"`
			State    string  `json:"state"`
			ActiveAt string  `json:"activeAt"`
			Value    float64 `json:"value"`
		} `json:"alerts"`
	} `json:"data"`
}

const response = `
{
    "status": "success",
    "data": {
        "alerts": [
            {
                "labels": {
                    "severity": "Low",
                    "ibm_resource_type": "member",
                    "ibm_scope": "a/ea727cda20c042788115c1cb01047f0a",
                    "ibm_ctype": "ibm-cloud",
                    "ibm_location": "eu-de",
                    "remote_write": "true",
                    "ibm_resource": "c-de6ed6dc-5245-43d8-82ca-a444ca4e35e8-m-0",
                    "ibm_resource_group_id": "2b0d12b9b0df49aa8d7a45c3fde9aec3",
                    "alertname": "LB_MS_POSTGRESQL_BLOCKS_DATABASE",
                    "ibm_service_instance": "de6ed6dc-5245-43d8-82ca-a444ca4e35e8",
                    "_sysdig_datasource": "prometheus_remote_write",
                    "ibm_service_name": "databases-for-postgresql",
                    "__name__": "ibm_databases_for_postgresql_blocks_hit_rate",
                    "ibm_resource_group_name": "ib-eude-agl-shs-dbcli--rg01-pro",
                    "_sysdig_custom_metric": "true",
                    "ibm_service_instance_name": "ib-eude-agl-shs-dbcli-intranet-dbpostgresql01-pro"
                },
                "annotations": {
                    "summary": "LB_MS_POSTGRESQL_BLOCKS_DATABASE is Triggered",
                    "description": "LB_MS_POSTGRESQL_BLOCKS_DATABASE"
                },
                "state": "firing",
                "activeAt": "2023-11-22T14:30:00.000+0000",
                "value": 242.5375
            },
            {
                "labels": {
                    "kube_pod_uid": "5ae22a62-59ce-4b71-a53a-7076fe988b22",
                    "kube_pod_label_helm_sh_chart": "itnow-php-7-3-3.1.0",
                    "kube_node_label_worker_application_type": "all",
                    "kube_namespace_name": "cxb-einave-pro",
                    "alertname": "LB_MS_IKS_POD_RESTART_APP_HIGH",
                    "kube_node_label_worker_pool": "worker",
                    "kube_replicaset_label_ssp_company": "cxb",
                    "kube_pod_label_ssp_company": "cxb",
                    "kube_node_label_topology_kubernetes_io_region": "eu-de",
                    "kube_replicaset_name": "einave-einavefront-pro-7b7d849bc4",
                    "kube_deployment_name": "einave-einavefront-pro",
                    "kube_cluster_name": "ib-eude-cxb-orc-k8isi--mks01-pro",
                    "kube_node_label_beta_kubernetes_io_os": "linux",
                    "agent_tag_cluster": "ib-eude-cxb-orc-k8isi--mks01-pro",
                    "kube_deployment_label_ssp_company": "cxb",
                    "kube_deployment_label_ssp_component": "einavefront",
                    "kube_deployment_label_logging_elk_stack": "platform",
                    "kube_pod_label_ssp_component": "einavefront",
                    "kube_namespace_uid": "db3c9bb7-0ab4-4fee-b8d9-a971d2f11087",
                    "_sysdig_datasource": "agent",
                    "kube_node_label_node_kubernetes_io_instance_type": "mx2.16x128",
                    "kube_node_uid": "758ad8ac-d4e3-40b0-ada3-768fab1474f0",
                    "kube_node_label_arch": "amd64",
                    "__name__": "changes(kube_pod_container_status_restarts_total{kube_namespace_label_type=\"app\", kube_cluster_name =~ \".*-pro\"}[15m]) >= 3",
                    "kube_pod_label_logging_elk_stack": "platform",
                    "kube_deployment_uid": "9504bc12-b633-4c49-a110-ed1004d59f3a",
                    "kube_namespace_label_deploymentArea": "intranet",
                    "kube_node_label_ibm_cloud_kubernetes_io_zone": "eu-de-2",
                    "kube_node_label_ibm_cloud_kubernetes_io_os": "UBUNTU_20_64",
                    "agent_tag_ibm_containers_kubernetes_cluster_id": "c99tr12f04ubqo7n5big",
                    "kube_node_name": "10.152.136.74",
                    "kube_namespace_label_ssp_squad": "squad_1",
                    "kube_replicaset_uid": "b9e67b47-b8ac-413b-bda4-bf776bbc496d",
                    "kube_namespace_label_istio_injection": "enabled",
                    "kube_node_label_kubernetes_io_hostname": "10.152.136.74",
                    "kube_pod_label_security_istio_io_tlsMode": "istio",
                    "kube_node_label_failure_domain_beta_kubernetes_io_region": "eu-de",
                    "kube_pod_container_name": "itnow-php-7-3",
                    "kube_cluster_id": "4eaae506-757b-4f76-a2d7-d3ef52e7c14f",
                    "kube_namespace_label_name": "einave",
                    "kube_replicaset_label_ssp_environment": "pro",
                    "kube_namespace_label_env": "pro",
                    "kube_namespace_label_company": "cxb",
                    "kube_pod_label_service_istio_io_canonical_name": "itnow-php-7-3",
                    "kube_replicaset_label_ssp_application": "einave",
                    "kube_namespace_label_type": "app",
                    "kube_node_label_failure_domain_beta_kubernetes_io_zone": "eu-de-2",
                    "kube_workload_name": "einave-einavefront-pro",
                    "kube_pod_name": "einave-einavefront-pro-7b7d849bc4-85tnw",
                    "kube_pod_label_service_istio_io_canonical_revision": "latest",
                    "kube_replicaset_label_ssp_component": "einavefront",
                    "severity": "High",
                    "kube_pod_internal_ip": "172.17.100.233",
                    "kube_node_label_node_role_kubernetes_io_worker": "true",
                    "kube_deployment_label_ssp_environment": "pro",
                    "kube_node_label_beta_kubernetes_io_instance_type": "mx2.16x128",
                    "kube_pod_label_ssp_application": "einave",
                    "kube_pod_label_topology_istio_io_network": "ibmcloudPublicPro-intranet",
                    "kube_workload_type": "deployment",
                    "kube_pod_label_ssp_environment": "pro",
                    "kube_replicaset_label_logging_elk_stack": "platform",
                    "kube_node_label_ibm_cloud_kubernetes_io_region": "eu-de",
                    "kube_deployment_label_ssp_application": "einave",
                    "kube_deployment_label_helm_sh_chart": "itnow-php-7-3-3.1.0",
                    "kube_node_label_kubernetes_io_os": "linux",
                    "kube_replicaset_label_helm_sh_chart": "itnow-php-7-3-3.1.0",
                    "kube_deployment_label_helm_sh_deploy_date": "2023-11-28_094426.058",
                    "kube_node_label_topology_kubernetes_io_zone": "eu-de-2"
                },
                "annotations": {
                    "summary": "LB_MS_IKS_POD_RESTART_APP_HIGH is Triggered",
                    "description": "LB_MS_IKS_POD_RESTART_APP_HIGH"
                },
                "state": "firing",
                "activeAt": "2023-11-28T10:03:00.000+0000",
                "value": 3.0
            }
        ]
    }
}`

func main() {
	// Check tutorial --> https://medium.com/swlh/unmarsahl-a-nested-json-in-go-c1a953aeda2
	var alertas SysdigAlertsResponse
	var exist bool

	err := json.Unmarshal([]byte(response), &alertas)
	if err != nil {
		log.Fatal("error unmarshaling json: ", err)
	}

	log.Println("Salida -->")
	for _, item := range alertas.Data.Alerts {
		tipo := item.Labels.Type
		if tipo != "" {
			log.Println("tipo --> ", tipo)
			exist = true
		}
		if exist {
			log.Println("En este caso, existe")
			tipoSimplificado := string(tipo)[0:3]
			tipoSimplificado = strings.ToLower(tipoSimplificado)
			log.Println("Resultado: ", tipoSimplificado)
		}

		if !exist {
			log.Println("La negativa del if")
			log.Println(item.Labels.Alertname)
		}
	}
}
