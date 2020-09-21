# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [common/policies.proto](#common/policies.proto)
    - [RecommendationPolicy](#containersai.common.RecommendationPolicy)
  
- [common/types.proto](#common/types.proto)
    - [ColumnType](#containersai.common.ColumnType)
    - [DataType](#containersai.common.DataType)
  
- [common/common.proto](#common/common.proto)
    - [Group](#containersai.common.Group)
    - [Query](#containersai.common.Query)
    - [QueryCondition](#containersai.common.QueryCondition)
    - [ReadRawdata](#containersai.common.ReadRawdata)
    - [Row](#containersai.common.Row)
    - [TimeRange](#containersai.common.TimeRange)
    - [WriteRawdata](#containersai.common.WriteRawdata)
  
    - [DatabaseType](#containersai.common.DatabaseType)
    - [QueryCondition.Order](#containersai.common.QueryCondition.Order)
    - [TimeRange.AggregateFunction](#containersai.common.TimeRange.AggregateFunction)
  
- [alameda_api/v1alpha1/datahub/predictions/services.proto](#alameda_api/v1alpha1/datahub/predictions/services.proto)
    - [CreateApplicationPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest)
    - [CreateClusterPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest)
    - [CreateControllerPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest)
    - [CreateNamespacePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest)
    - [CreateNodePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest)
    - [CreatePodPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest)
    - [CreatePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest)
    - [ListApplicationPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest)
    - [ListApplicationPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse)
    - [ListClusterPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest)
    - [ListClusterPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse)
    - [ListControllerPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest)
    - [ListControllerPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse)
    - [ListNamespacePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest)
    - [ListNamespacePredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse)
    - [ListNodePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest)
    - [ListNodePredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse)
    - [ListPodPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest)
    - [ListPodPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse)
    - [ListPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest)
    - [ListPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse)
  
- [alameda_api/v1alpha1/datahub/predictions/types.proto](#alameda_api/v1alpha1/datahub/predictions/types.proto)
    - [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData)
    - [Prediction](#containersai.alameda.v1alpha1.datahub.predictions.Prediction)
    - [PredictionData](#containersai.alameda.v1alpha1.datahub.predictions.PredictionData)
    - [Sample](#containersai.alameda.v1alpha1.datahub.predictions.Sample)
  
- [alameda_api/v1alpha1/datahub/predictions/predictions.proto](#alameda_api/v1alpha1/datahub/predictions/predictions.proto)
    - [ApplicationPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction)
    - [ClusterPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction)
    - [ContainerPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction)
    - [ControllerPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction)
    - [NamespacePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction)
    - [NodePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NodePrediction)
    - [PodPrediction](#containersai.alameda.v1alpha1.datahub.predictions.PodPrediction)
    - [WritePrediction](#containersai.alameda.v1alpha1.datahub.predictions.WritePrediction)
  
- [alameda_api/v1alpha1/datahub/events/services.proto](#alameda_api/v1alpha1/datahub/events/services.proto)
    - [CreateEventsRequest](#containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest)
    - [ListEventsRequest](#containersai.alameda.v1alpha1.datahub.events.ListEventsRequest)
    - [ListEventsResponse](#containersai.alameda.v1alpha1.datahub.events.ListEventsResponse)
  
- [alameda_api/v1alpha1/datahub/events/types.proto](#alameda_api/v1alpha1/datahub/events/types.proto)
    - [EventSource](#containersai.alameda.v1alpha1.datahub.events.EventSource)
    - [K8SObjectReference](#containersai.alameda.v1alpha1.datahub.events.K8SObjectReference)
  
    - [EventLevel](#containersai.alameda.v1alpha1.datahub.events.EventLevel)
    - [EventType](#containersai.alameda.v1alpha1.datahub.events.EventType)
    - [EventVersion](#containersai.alameda.v1alpha1.datahub.events.EventVersion)
  
- [alameda_api/v1alpha1/datahub/events/events.proto](#alameda_api/v1alpha1/datahub/events/events.proto)
    - [Event](#containersai.alameda.v1alpha1.datahub.events.Event)
  
- [alameda_api/v1alpha1/datahub/metrics/services.proto](#alameda_api/v1alpha1/datahub/metrics/services.proto)
    - [CreateApplicationMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest)
    - [CreateClusterMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest)
    - [CreateControllerMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest)
    - [CreateMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest)
    - [CreateNamespaceMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest)
    - [CreateNodeMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest)
    - [CreatePodMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest)
    - [ListApplicationMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest)
    - [ListApplicationMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse)
    - [ListClusterMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest)
    - [ListClusterMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse)
    - [ListControllerMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest)
    - [ListControllerMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse)
    - [ListMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest)
    - [ListMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse)
    - [ListNamespaceMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest)
    - [ListNamespaceMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse)
    - [ListNodeMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest)
    - [ListNodeMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse)
    - [ListPodMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest)
    - [ListPodMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse)
  
- [alameda_api/v1alpha1/datahub/metrics/types.proto](#alameda_api/v1alpha1/datahub/metrics/types.proto)
    - [Metric](#containersai.alameda.v1alpha1.datahub.metrics.Metric)
    - [MetricData](#containersai.alameda.v1alpha1.datahub.metrics.MetricData)
  
- [alameda_api/v1alpha1/datahub/metrics/metrics.proto](#alameda_api/v1alpha1/datahub/metrics/metrics.proto)
    - [ApplicationMetric](#containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric)
    - [ClusterMetric](#containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric)
    - [ContainerMetric](#containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric)
    - [ControllerMetric](#containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric)
    - [NamespaceMetric](#containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric)
    - [NodeMetric](#containersai.alameda.v1alpha1.datahub.metrics.NodeMetric)
    - [PodMetric](#containersai.alameda.v1alpha1.datahub.metrics.PodMetric)
    - [WriteMetric](#containersai.alameda.v1alpha1.datahub.metrics.WriteMetric)
  
- [alameda_api/v1alpha1/datahub/applications/services.proto](#alameda_api/v1alpha1/datahub/applications/services.proto)
    - [CreateApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest)
    - [DeleteApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest)
    - [ListApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest)
    - [ListApplicationsResponse](#containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse)
  
- [alameda_api/v1alpha1/datahub/applications/types.proto](#alameda_api/v1alpha1/datahub/applications/types.proto)
    - [Application](#containersai.alameda.v1alpha1.datahub.applications.Application)
    - [ApplicationData](#containersai.alameda.v1alpha1.datahub.applications.ApplicationData)
  
- [alameda_api/v1alpha1/datahub/applications/applications.proto](#alameda_api/v1alpha1/datahub/applications/applications.proto)
    - [DeleteApplication](#containersai.alameda.v1alpha1.datahub.applications.DeleteApplication)
    - [ReadApplication](#containersai.alameda.v1alpha1.datahub.applications.ReadApplication)
    - [WriteApplication](#containersai.alameda.v1alpha1.datahub.applications.WriteApplication)
  
- [alameda_api/v1alpha1/datahub/recommendations/services.proto](#alameda_api/v1alpha1/datahub/recommendations/services.proto)
    - [CreateApplicationRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest)
    - [CreateClusterRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest)
    - [CreateControllerRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest)
    - [CreateNamespaceRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest)
    - [CreateNodeRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest)
    - [CreatePodRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest)
    - [CreateRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest)
    - [ListApplicationRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest)
    - [ListApplicationRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse)
    - [ListClusterRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest)
    - [ListClusterRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse)
    - [ListControllerRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest)
    - [ListControllerRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse)
    - [ListNamespaceRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest)
    - [ListNamespaceRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse)
    - [ListNodeRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest)
    - [ListNodeRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse)
    - [ListPodRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest)
    - [ListPodRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse)
    - [ListRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest)
    - [ListRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse)
  
- [alameda_api/v1alpha1/datahub/recommendations/types.proto](#alameda_api/v1alpha1/datahub/recommendations/types.proto)
    - [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec)
    - [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s)
    - [Recommendation](#containersai.alameda.v1alpha1.datahub.recommendations.Recommendation)
    - [RecommendationData](#containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData)
  
    - [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType)
  
- [alameda_api/v1alpha1/datahub/recommendations/recommendations.proto](#alameda_api/v1alpha1/datahub/recommendations/recommendations.proto)
    - [ApplicationRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ApplicationRecommendation)
    - [ClusterRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ClusterRecommendation)
    - [ContainerRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ContainerRecommendation)
    - [ControllerRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendation)
    - [NamespaceRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NamespaceRecommendation)
    - [NodeRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NodeRecommendation)
    - [PodRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.PodRecommendation)
    - [WriteRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.WriteRecommendation)
  
- [alameda_api/v1alpha1/datahub/common/types.proto](#alameda_api/v1alpha1/datahub/common/types.proto)
    - [ColumnType](#containersai.alameda.v1alpha1.datahub.common.ColumnType)
    - [DataType](#containersai.alameda.v1alpha1.datahub.common.DataType)
    - [DatabaseType](#containersai.alameda.v1alpha1.datahub.common.DatabaseType)
    - [FunctionType](#containersai.alameda.v1alpha1.datahub.common.FunctionType)
    - [ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary)
    - [ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota)
  
- [alameda_api/v1alpha1/datahub/common/metrics.proto](#alameda_api/v1alpha1/datahub/common/metrics.proto)
    - [MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData)
    - [Sample](#containersai.alameda.v1alpha1.datahub.common.Sample)
  
    - [MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType)
    - [ResourceName](#containersai.alameda.v1alpha1.datahub.common.ResourceName)
  
- [alameda_api/v1alpha1/datahub/common/queries.proto](#alameda_api/v1alpha1/datahub/common/queries.proto)
    - [Condition](#containersai.alameda.v1alpha1.datahub.common.Condition)
    - [Function](#containersai.alameda.v1alpha1.datahub.common.Function)
    - [Into](#containersai.alameda.v1alpha1.datahub.common.Into)
    - [QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition)
    - [TimeRange](#containersai.alameda.v1alpha1.datahub.common.TimeRange)
  
    - [QueryCondition.Order](#containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order)
    - [TimeRange.AggregateFunction](#containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction)
  
- [alameda_api/v1alpha1/datahub/common/rawdata.proto](#alameda_api/v1alpha1/datahub/common/rawdata.proto)
    - [Group](#containersai.alameda.v1alpha1.datahub.common.Group)
    - [ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData)
    - [Row](#containersai.alameda.v1alpha1.datahub.common.Row)
    - [WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData)
  
- [alameda_api/v1alpha1/datahub/weavescope/services.proto](#alameda_api/v1alpha1/datahub/weavescope/services.proto)
    - [ListWeaveScopeContainersRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest)
    - [ListWeaveScopeHostsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest)
    - [ListWeaveScopePodsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest)
    - [WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)
  
- [alameda_api/v1alpha1/datahub/rawdata/services.proto](#alameda_api/v1alpha1/datahub/rawdata/services.proto)
    - [ReadRawdataRequest](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest)
    - [ReadRawdataResponse](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse)
    - [WriteRawdataRequest](#containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest)
  
- [alameda_api/v1alpha1/datahub/rawdata/types.proto](#alameda_api/v1alpha1/datahub/rawdata/types.proto)
    - [Query](#containersai.alameda.v1alpha1.datahub.rawdata.Query)
  
- [alameda_api/v1alpha1/datahub/rawdata/rawdata.proto](#alameda_api/v1alpha1/datahub/rawdata/rawdata.proto)
    - [ReadRawdata](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata)
    - [WriteRawdata](#containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata)
  
- [alameda_api/v1alpha1/datahub/licenses/services.proto](#alameda_api/v1alpha1/datahub/licenses/services.proto)
    - [GetLicenseResponse](#containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse)
  
- [alameda_api/v1alpha1/datahub/licenses/licenses.proto](#alameda_api/v1alpha1/datahub/licenses/licenses.proto)
    - [License](#containersai.alameda.v1alpha1.datahub.licenses.License)
  
- [alameda_api/v1alpha1/datahub/scores/services.proto](#alameda_api/v1alpha1/datahub/scores/services.proto)
    - [CreateSimulatedSchedulingScoresRequest](#containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest)
    - [ListSimulatedSchedulingScoresRequest](#containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest)
    - [ListSimulatedSchedulingScoresResponse](#containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse)
  
- [alameda_api/v1alpha1/datahub/scores/scores.proto](#alameda_api/v1alpha1/datahub/scores/scores.proto)
    - [SimulatedSchedulingScore](#containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScore)
  
- [alameda_api/v1alpha1/datahub/data/services.proto](#alameda_api/v1alpha1/datahub/data/services.proto)
    - [DeleteDataRequest](#containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest)
    - [ReadDataRequest](#containersai.alameda.v1alpha1.datahub.data.ReadDataRequest)
    - [ReadDataResponse](#containersai.alameda.v1alpha1.datahub.data.ReadDataResponse)
    - [WriteDataRequest](#containersai.alameda.v1alpha1.datahub.data.WriteDataRequest)
    - [WriteMetaRequest](#containersai.alameda.v1alpha1.datahub.data.WriteMetaRequest)
  
- [alameda_api/v1alpha1/datahub/data/types.proto](#alameda_api/v1alpha1/datahub/data/types.proto)
    - [Data](#containersai.alameda.v1alpha1.datahub.data.Data)
    - [Rawdata](#containersai.alameda.v1alpha1.datahub.data.Rawdata)
  
- [alameda_api/v1alpha1/datahub/data/data.proto](#alameda_api/v1alpha1/datahub/data/data.proto)
    - [DeleteData](#containersai.alameda.v1alpha1.datahub.data.DeleteData)
    - [ReadData](#containersai.alameda.v1alpha1.datahub.data.ReadData)
    - [WriteData](#containersai.alameda.v1alpha1.datahub.data.WriteData)
    - [WriteMeta](#containersai.alameda.v1alpha1.datahub.data.WriteMeta)
  
- [alameda_api/v1alpha1/datahub/resources/services.proto](#alameda_api/v1alpha1/datahub/resources/services.proto)
    - [CreateApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest)
    - [CreateClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest)
    - [CreateControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest)
    - [CreateNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest)
    - [CreateNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest)
    - [CreatePodsRequest](#containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest)
    - [DeleteApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest)
    - [DeleteClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest)
    - [DeleteControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest)
    - [DeleteNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest)
    - [DeleteNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest)
    - [DeletePodsRequest](#containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest)
    - [ListApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest)
    - [ListApplicationsResponse](#containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse)
    - [ListClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest)
    - [ListClustersResponse](#containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse)
    - [ListControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest)
    - [ListControllersResponse](#containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse)
    - [ListNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest)
    - [ListNamespacesResponse](#containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse)
    - [ListNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest)
    - [ListNodesResponse](#containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse)
    - [ListPodsRequest](#containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest)
    - [ListPodsResponse](#containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse)
  
- [alameda_api/v1alpha1/datahub/resources/metadata.proto](#alameda_api/v1alpha1/datahub/resources/metadata.proto)
    - [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta)
    - [OwnerReference](#containersai.alameda.v1alpha1.datahub.resources.OwnerReference)
  
    - [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind)
    - [ScalingTool](#containersai.alameda.v1alpha1.datahub.resources.ScalingTool)
  
- [alameda_api/v1alpha1/datahub/resources/policies.proto](#alameda_api/v1alpha1/datahub/resources/policies.proto)
    - [AssignPodPolicy](#containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy)
    - [NodePriority](#containersai.alameda.v1alpha1.datahub.resources.NodePriority)
    - [Selector](#containersai.alameda.v1alpha1.datahub.resources.Selector)
    - [Selector.SelectorEntry](#containersai.alameda.v1alpha1.datahub.resources.Selector.SelectorEntry)
  
    - [RecommendationPolicy](#containersai.alameda.v1alpha1.datahub.resources.RecommendationPolicy)
  
- [alameda_api/v1alpha1/datahub/resources/types.proto](#alameda_api/v1alpha1/datahub/resources/types.proto)
    - [AlamedaApplicationSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaApplicationSpec)
    - [AlamedaControllerSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaControllerSpec)
    - [AlamedaNodeSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaNodeSpec)
    - [AlamedaPodSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaPodSpec)
    - [Capacity](#containersai.alameda.v1alpha1.datahub.resources.Capacity)
    - [Provider](#containersai.alameda.v1alpha1.datahub.resources.Provider)
    - [ResourceRequirements](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements)
    - [ResourceRequirements.LimitsEntry](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.LimitsEntry)
    - [ResourceRequirements.RequestsEntry](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.RequestsEntry)
  
- [alameda_api/v1alpha1/datahub/resources/status.proto](#alameda_api/v1alpha1/datahub/resources/status.proto)
    - [ContainerState](#containersai.alameda.v1alpha1.datahub.resources.ContainerState)
    - [ContainerStateRunning](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateRunning)
    - [ContainerStateTerminated](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateTerminated)
    - [ContainerStateWaiting](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateWaiting)
    - [ContainerStatus](#containersai.alameda.v1alpha1.datahub.resources.ContainerStatus)
    - [PodStatus](#containersai.alameda.v1alpha1.datahub.resources.PodStatus)
  
    - [PodPhase](#containersai.alameda.v1alpha1.datahub.resources.PodPhase)
  
- [alameda_api/v1alpha1/datahub/resources/resources.proto](#alameda_api/v1alpha1/datahub/resources/resources.proto)
    - [Application](#containersai.alameda.v1alpha1.datahub.resources.Application)
    - [Cluster](#containersai.alameda.v1alpha1.datahub.resources.Cluster)
    - [Container](#containersai.alameda.v1alpha1.datahub.resources.Container)
    - [Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller)
    - [Namespace](#containersai.alameda.v1alpha1.datahub.resources.Namespace)
    - [Node](#containersai.alameda.v1alpha1.datahub.resources.Node)
    - [Pod](#containersai.alameda.v1alpha1.datahub.resources.Pod)
  
- [alameda_api/v1alpha1/datahub/keycodes/services.proto](#alameda_api/v1alpha1/datahub/keycodes/services.proto)
    - [ActivateRegistrationDataRequest](#containersai.alameda.v1alpha1.datahub.keycodes.ActivateRegistrationDataRequest)
    - [AddKeycodeRequest](#containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeRequest)
    - [AddKeycodeResponse](#containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeResponse)
    - [DeleteKeycodeRequest](#containersai.alameda.v1alpha1.datahub.keycodes.DeleteKeycodeRequest)
    - [GenerateRegistrationDataResponse](#containersai.alameda.v1alpha1.datahub.keycodes.GenerateRegistrationDataResponse)
    - [ListKeycodesRequest](#containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesRequest)
    - [ListKeycodesResponse](#containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesResponse)
  
- [alameda_api/v1alpha1/datahub/keycodes/keycodes.proto](#alameda_api/v1alpha1/datahub/keycodes/keycodes.proto)
    - [Keycode](#containersai.alameda.v1alpha1.datahub.keycodes.Keycode)
  
- [alameda_api/v1alpha1/datahub/keycodes/types.proto](#alameda_api/v1alpha1/datahub/keycodes/types.proto)
    - [Capacity](#containersai.alameda.v1alpha1.datahub.keycodes.Capacity)
    - [Functionality](#containersai.alameda.v1alpha1.datahub.keycodes.Functionality)
    - [Retention](#containersai.alameda.v1alpha1.datahub.keycodes.Retention)
    - [ServiceAgreement](#containersai.alameda.v1alpha1.datahub.keycodes.ServiceAgreement)
  
- [alameda_api/v1alpha1/datahub/server.proto](#alameda_api/v1alpha1/datahub/server.proto)
    - [DatahubService](#containersai.alameda.v1alpha1.datahub.DatahubService)
  
- [alameda_api/v1alpha1/datahub/plannings/services.proto](#alameda_api/v1alpha1/datahub/plannings/services.proto)
    - [CreateApplicationPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest)
    - [CreateClusterPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest)
    - [CreateControllerPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest)
    - [CreateNamespacePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest)
    - [CreateNodePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest)
    - [CreatePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest)
    - [CreatePodPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest)
    - [ListApplicationPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest)
    - [ListApplicationPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse)
    - [ListClusterPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest)
    - [ListClusterPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse)
    - [ListControllerPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest)
    - [ListControllerPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse)
    - [ListNamespacePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest)
    - [ListNamespacePlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse)
    - [ListNodePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest)
    - [ListNodePlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse)
    - [ListPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest)
    - [ListPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse)
    - [ListPodPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest)
    - [ListPodPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse)
  
- [alameda_api/v1alpha1/datahub/plannings/types.proto](#alameda_api/v1alpha1/datahub/plannings/types.proto)
    - [ControllerPlanningSpec](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningSpec)
    - [ControllerPlanningSpecK8s](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningSpecK8s)
    - [PlanningData](#containersai.alameda.v1alpha1.datahub.plannings.PlanningData)
    - [RawPlanning](#containersai.alameda.v1alpha1.datahub.plannings.RawPlanning)
  
    - [ControllerPlanningType](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningType)
    - [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType)
  
- [alameda_api/v1alpha1/datahub/plannings/plannings.proto](#alameda_api/v1alpha1/datahub/plannings/plannings.proto)
    - [ApplicationPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning)
    - [ClusterPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning)
    - [ContainerPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning)
    - [ControllerPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning)
    - [NamespacePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning)
    - [NodePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NodePlanning)
    - [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning)
    - [PodPlanning](#containersai.alameda.v1alpha1.datahub.plannings.PodPlanning)
    - [WritePlanning](#containersai.alameda.v1alpha1.datahub.plannings.WritePlanning)
  
- [alameda_api/v1alpha1/datahub/gpu/services.proto](#alameda_api/v1alpha1/datahub/gpu/services.proto)
    - [CreateGpuPredictionsRequest](#containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest)
    - [ListGpuMetricsRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest)
    - [ListGpuMetricsResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse)
    - [ListGpuPredictionsRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest)
    - [ListGpuPredictionsResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse)
    - [ListGpusRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest)
    - [ListGpusResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse)
  
- [alameda_api/v1alpha1/datahub/gpu/types.proto](#alameda_api/v1alpha1/datahub/gpu/types.proto)
    - [GpuMetadata](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata)
    - [GpuSpec](#containersai.alameda.v1alpha1.datahub.gpu.GpuSpec)
  
- [alameda_api/v1alpha1/datahub/gpu/gpu.proto](#alameda_api/v1alpha1/datahub/gpu/gpu.proto)
    - [Gpu](#containersai.alameda.v1alpha1.datahub.gpu.Gpu)
    - [GpuMetric](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetric)
    - [GpuPrediction](#containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction)
  
- [alameda_api/v1alpha1/datahub/schemas/services.proto](#alameda_api/v1alpha1/datahub/schemas/services.proto)
    - [CreateSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest)
    - [DeleteSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest)
    - [ListSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest)
    - [ListSchemasResponse](#containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse)
  
- [alameda_api/v1alpha1/datahub/schemas/types.proto](#alameda_api/v1alpha1/datahub/schemas/types.proto)
    - [Column](#containersai.alameda.v1alpha1.datahub.schemas.Column)
    - [Measurement](#containersai.alameda.v1alpha1.datahub.schemas.Measurement)
    - [SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta)
  
    - [Scope](#containersai.alameda.v1alpha1.datahub.schemas.Scope)
  
- [alameda_api/v1alpha1/datahub/schemas/schemas.proto](#alameda_api/v1alpha1/datahub/schemas/schemas.proto)
    - [Schema](#containersai.alameda.v1alpha1.datahub.schemas.Schema)
  
- [datahub/recommendation/v1alpha2/recommendation.proto](#datahub/recommendation/v1alpha2/recommendation.proto)
    - [AssignPodPolicy](#containersai.datahub.recommendation.v1alpha2.AssignPodPolicy)
    - [ContainerRecommendation](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation)
    - [ContainerRecommendation.InitialLimitRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialLimitRecommendationsEntry)
    - [ContainerRecommendation.InitialRequestRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialRequestRecommendationsEntry)
    - [ContainerRecommendation.LimitRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.LimitRecommendationsEntry)
    - [ContainerRecommendation.RequestRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.RequestRecommendationsEntry)
    - [PodRecommendation](#containersai.datahub.recommendation.v1alpha2.PodRecommendation)
  
    - [RecommendationPolicy](#containersai.datahub.recommendation.v1alpha2.RecommendationPolicy)
  
- [datahub/score/v1alpha2/score.proto](#datahub/score/v1alpha2/score.proto)
    - [SimulatedSchedulingScore](#containersai.datahub.score.v1alpha2.SimulatedSchedulingScore)
  
- [datahub/v1alpha2/datahub.proto](#datahub/v1alpha2/datahub.proto)
    - [CreateNodePredictionsRequest](#containersai.datahub.v1alpha2.CreateNodePredictionsRequest)
    - [CreateNodesRequest](#containersai.datahub.v1alpha2.CreateNodesRequest)
    - [CreatePodPredictionsRequest](#containersai.datahub.v1alpha2.CreatePodPredictionsRequest)
    - [CreatePodRecommendationsRequest](#containersai.datahub.v1alpha2.CreatePodRecommendationsRequest)
    - [CreatePodsRequest](#containersai.datahub.v1alpha2.CreatePodsRequest)
    - [CreateSimulatedSchedulingScoresRequest](#containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest)
    - [DeleteNodesRequest](#containersai.datahub.v1alpha2.DeleteNodesRequest)
    - [DeletePodsRequest](#containersai.datahub.v1alpha2.DeletePodsRequest)
    - [ListNodeMetricsRequest](#containersai.datahub.v1alpha2.ListNodeMetricsRequest)
    - [ListNodeMetricsResponse](#containersai.datahub.v1alpha2.ListNodeMetricsResponse)
    - [ListNodePredictionsRequest](#containersai.datahub.v1alpha2.ListNodePredictionsRequest)
    - [ListNodePredictionsResponse](#containersai.datahub.v1alpha2.ListNodePredictionsResponse)
    - [ListNodesRequest](#containersai.datahub.v1alpha2.ListNodesRequest)
    - [ListNodesResponse](#containersai.datahub.v1alpha2.ListNodesResponse)
    - [ListPodMetricsRequest](#containersai.datahub.v1alpha2.ListPodMetricsRequest)
    - [ListPodMetricsResponse](#containersai.datahub.v1alpha2.ListPodMetricsResponse)
    - [ListPodPredictionsRequest](#containersai.datahub.v1alpha2.ListPodPredictionsRequest)
    - [ListPodPredictionsResponse](#containersai.datahub.v1alpha2.ListPodPredictionsResponse)
    - [ListPodRecommendationsRequest](#containersai.datahub.v1alpha2.ListPodRecommendationsRequest)
    - [ListPodRecommendationsResponse](#containersai.datahub.v1alpha2.ListPodRecommendationsResponse)
    - [ListPodsRequest](#containersai.datahub.v1alpha2.ListPodsRequest)
    - [ListPodsResponse](#containersai.datahub.v1alpha2.ListPodsResponse)
    - [ListSimulatedSchedulingScoresRequest](#containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest)
    - [ListSimulatedSchedulingScoresResponse](#containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse)
    - [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition)
    - [UpdateNodesRequest](#containersai.datahub.v1alpha2.UpdateNodesRequest)
    - [UpdateNodesRequest.UpdatedNode](#containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode)
    - [UpdateNodesRequest.UpdatedNode.IsPredictedWrap](#containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode.IsPredictedWrap)
    - [UpdatePodsRequest](#containersai.datahub.v1alpha2.UpdatePodsRequest)
    - [UpdatePodsRequest.UpdatedPod](#containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod)
    - [UpdatePodsRequest.UpdatedPod.IsPredictedWrap](#containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod.IsPredictedWrap)
  
    - [QueryCondition.Order](#containersai.datahub.v1alpha2.QueryCondition.Order)
  
    - [DatahubService](#containersai.datahub.v1alpha2.DatahubService)
  
- [datahub/metric/v1alpha2/metric.proto](#datahub/metric/v1alpha2/metric.proto)
    - [ContainerMetric](#containersai.datahub.metric.v1alpha2.ContainerMetric)
    - [ContainerMetric.MetricDataEntry](#containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry)
    - [MetricData](#containersai.datahub.metric.v1alpha2.MetricData)
    - [NodeMetric](#containersai.datahub.metric.v1alpha2.NodeMetric)
    - [NodeMetric.MetricDataEntry](#containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry)
    - [PodMetric](#containersai.datahub.metric.v1alpha2.PodMetric)
    - [Sample](#containersai.datahub.metric.v1alpha2.Sample)
    - [TimeRange](#containersai.datahub.metric.v1alpha2.TimeRange)
  
    - [MetricType](#containersai.datahub.metric.v1alpha2.MetricType)
  
- [datahub/resource/pod/assign/v1alpha2/assign.proto](#datahub/resource/pod/assign/v1alpha2/assign.proto)
    - [NodePriority](#containersai.datahub.resource.pod.assign.v1alpha2.NodePriority)
    - [Selector](#containersai.datahub.resource.pod.assign.v1alpha2.Selector)
    - [Selector.SelectorEntry](#containersai.datahub.resource.pod.assign.v1alpha2.Selector.SelectorEntry)
  
- [datahub/resource/v1alpha2/resource.proto](#datahub/resource/v1alpha2/resource.proto)
    - [Container](#containersai.datahub.resource.v1alpha2.Container)
    - [Container.LimitResourceEntry](#containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry)
    - [Container.RequestResourceEntry](#containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry)
    - [Node](#containersai.datahub.resource.v1alpha2.Node)
    - [Pod](#containersai.datahub.resource.v1alpha2.Pod)
  
- [datahub/resource/metadata/v1alpha2/metadata.proto](#datahub/resource/metadata/v1alpha2/metadata.proto)
    - [NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName)
  
- [datahub/prediction/v1alpha2/prediction.proto](#datahub/prediction/v1alpha2/prediction.proto)
    - [ContainerPrediction](#containersai.datahub.prediction.v1alpha2.ContainerPrediction)
    - [ContainerPrediction.PredictedRawDataEntry](#containersai.datahub.prediction.v1alpha2.ContainerPrediction.PredictedRawDataEntry)
    - [NodePrediction](#containersai.datahub.prediction.v1alpha2.NodePrediction)
    - [NodePrediction.PredictedRawDataEntry](#containersai.datahub.prediction.v1alpha2.NodePrediction.PredictedRawDataEntry)
    - [PodPrediction](#containersai.datahub.prediction.v1alpha2.PodPrediction)
  
- [Scalar Value Types](#scalar-value-types)



<a name="common/policies.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/policies.proto


 


<a name="containersai.common.RecommendationPolicy"></a>

### RecommendationPolicy
Recommendation policy. A policy may be either stable or compact.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RECOMMENDATION_POLICY_UNDEFINED | 0 |  |
| STABLE | 1 |  |
| COMPACT | 2 |  |


 

 

 



<a name="common/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/types.proto


 


<a name="containersai.common.ColumnType"></a>

### ColumnType
Represents the field type of a record which queried from datahub.

| Name | Number | Description |
| ---- | ------ | ----------- |
| COLUMNTYPE_UDEFINED | 0 |  |
| COLUMNTYPE_TAG | 1 |  |
| COLUMNTYPE_FIELD | 2 |  |



<a name="containersai.common.DataType"></a>

### DataType
Represents the datahub specified data type.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DATATYPE_UNDEFINED | 0 |  |
| DATATYPE_BOOL | 1 |  |
| DATATYPE_INT | 2 |  |
| DATATYPE_INT8 | 3 |  |
| DATATYPE_INT16 | 4 |  |
| DATATYPE_INT32 | 5 |  |
| DATATYPE_INT64 | 6 |  |
| DATATYPE_UINT | 7 |  |
| DATATYPE_UINT8 | 8 |  |
| DATATYPE_UINT16 | 9 |  |
| DATATYPE_UINT32 | 10 |  |
| DATATYPE_UTIN64 | 11 |  |
| DATATYPE_FLOAT32 | 12 |  |
| DATATYPE_FLOAT64 | 13 |  |
| DATATYPE_STRING | 14 |  |


 

 

 



<a name="common/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/common.proto



<a name="containersai.common.Group"></a>

### Group
Represents a dataset which are collected that have the same attributes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rows | [Row](#containersai.common.Row) | repeated |  |






<a name="containersai.common.Query"></a>

### Query
Represents a general datahub query request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| table | [string](#string) |  |  |
| expression | [string](#string) |  |  |
| condition | [QueryCondition](#containersai.common.QueryCondition) |  |  |






<a name="containersai.common.QueryCondition"></a>

### QueryCondition
Represents a datahub query request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time_range | [TimeRange](#containersai.common.TimeRange) |  |  |
| order | [QueryCondition.Order](#containersai.common.QueryCondition.Order) |  |  |
| where_clause | [string](#string) |  |  |
| selects | [string](#string) | repeated |  |
| groups | [string](#string) | repeated |  |
| limit | [uint64](#uint64) |  |  |






<a name="containersai.common.ReadRawdata"></a>

### ReadRawdata
Represents a rawdata whcih is read from datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [Query](#containersai.common.Query) |  |  |
| columns | [string](#string) | repeated |  |
| groups | [Group](#containersai.common.Group) | repeated |  |
| rawdata | [string](#string) |  |  |






<a name="containersai.common.Row"></a>

### Row
Represents a record of data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| values | [string](#string) | repeated |  |






<a name="containersai.common.TimeRange"></a>

### TimeRange
Represents a time range definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| timeout | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| step | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| aggregate_function | [TimeRange.AggregateFunction](#containersai.common.TimeRange.AggregateFunction) |  |  |
| apply_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="containersai.common.WriteRawdata"></a>

### WriteRawdata
Represents a rawdata which will be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| table | [string](#string) |  |  |
| columns | [string](#string) | repeated |  |
| rows | [Row](#containersai.common.Row) | repeated |  |
| column_types | [ColumnType](#containersai.common.ColumnType) | repeated |  |
| data_types | [DataType](#containersai.common.DataType) | repeated |  |





 


<a name="containersai.common.DatabaseType"></a>

### DatabaseType
Represents a specified database whcih is to query.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 |  |
| INFLUXDB | 1 |  |
| PROMETHEUS | 2 |  |



<a name="containersai.common.QueryCondition.Order"></a>

### QueryCondition.Order


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ASC | 1 |  |
| DESC | 2 |  |



<a name="containersai.common.TimeRange.AggregateFunction"></a>

### TimeRange.AggregateFunction


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| MAX | 1 |  |
| AVG | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/predictions/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/predictions/services.proto



<a name="containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest"></a>

### CreateApplicationPredictionsRequest
Represents a request for creating predictions of alameda scalers&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_predictions | [ApplicationPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest"></a>

### CreateClusterPredictionsRequest
Represents a request for creating predictions clusters&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_predictions | [ClusterPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest"></a>

### CreateControllerPredictionsRequest
Represents a request for creating predictions of controllers&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| controller_predictions | [ControllerPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest"></a>

### CreateNamespacePredictionsRequest
Represents a request for creating predictions of namespaces&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace_predictions | [NamespacePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest"></a>

### CreateNodePredictionsRequest
Represents a request for creating predictions nodes&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_predictions | [NodePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NodePrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest"></a>

### CreatePodPredictionsRequest
Represents a request for creating predictions of containers&#39; metric data belonging to a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_predictions | [PodPrediction](#containersai.alameda.v1alpha1.datahub.predictions.PodPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest"></a>

### CreatePredictionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| predictions | [WritePrediction](#containersai.alameda.v1alpha1.datahub.predictions.WritePrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest"></a>

### ListApplicationPredictionsRequest
Represents a request for listing predictions of alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse"></a>

### ListApplicationPredictionsResponse
Represents a response for a listing predictions of alameda scalers request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| application_predictions | [ApplicationPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest"></a>

### ListClusterPredictionsRequest
Represents a request for listing predictions of clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse"></a>

### ListClusterPredictionsResponse
Represents a response for a listing predictions of clusters request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| cluster_predictions | [ClusterPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest"></a>

### ListControllerPredictionsRequest
Represents a request for listing predictions of controllers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse"></a>

### ListControllerPredictionsResponse
Represents a response for a listing predictions of controllers request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| controller_predictions | [ControllerPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest"></a>

### ListNamespacePredictionsRequest
Represents a request for listing predictions of namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse"></a>

### ListNamespacePredictionsResponse
Represents a response for a listing predictions of namespaces request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| namespace_predictions | [NamespacePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest"></a>

### ListNodePredictionsRequest
Represents a request for listing predictions of nodes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse"></a>

### ListNodePredictionsResponse
Represents a response for a listing predictions of nodes request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_predictions | [NodePrediction](#containersai.alameda.v1alpha1.datahub.predictions.NodePrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest"></a>

### ListPodPredictionsRequest
Represents a request for listing predictions of pods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| granularity | [int64](#int64) |  |  |
| fill_days | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse"></a>

### ListPodPredictionsResponse
Represents a response for a listing predictions of pods request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_predictions | [PodPrediction](#containersai.alameda.v1alpha1.datahub.predictions.PodPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest"></a>

### ListPredictionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse"></a>

### ListPredictionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| predictions | [Prediction](#containersai.alameda.v1alpha1.datahub.predictions.Prediction) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/predictions/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/predictions/types.proto



<a name="containersai.alameda.v1alpha1.datahub.predictions.MetricData"></a>

### MetricData
Represents a piece of metreic data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| data | [Sample](#containersai.alameda.v1alpha1.datahub.predictions.Sample) | repeated | data can be time series or non-time series |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.Prediction"></a>

### Prediction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| prediction_data | [PredictionData](#containersai.alameda.v1alpha1.datahub.predictions.PredictionData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.PredictionData"></a>

### PredictionData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| read_data | [containersai.alameda.v1alpha1.datahub.common.ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.Sample"></a>

### Sample
Represents a data point of time-series metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| num_value | [string](#string) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/predictions/predictions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/predictions/predictions.proto



<a name="containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction"></a>

### ApplicationPrediction
Represents a list of predicted metrics data of a alameda scaler.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction"></a>

### ClusterPrediction
Represents a list of predicted metric data of a cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction"></a>

### ContainerPrediction
Represents a list of predicted metric data of a container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction"></a>

### ControllerPrediction
Represents a list of predicted metrics data of a controller.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction"></a>

### NamespacePrediction
Represents a list of predicted metrics data of a namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.NodePrediction"></a>

### NodePrediction
Represents a list of predicted metric data of a node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| is_scheduled | [bool](#bool) |  |  |
| predicted_raw_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.PodPrediction"></a>

### PodPrediction
Represents a list of predicted metrics data of a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| container_predictions | [ContainerPrediction](#containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.predictions.WritePrediction"></a>

### WritePrediction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| write_data | [containersai.alameda.v1alpha1.datahub.common.WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/events/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/events/services.proto



<a name="containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest"></a>

### CreateEventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#containersai.alameda.v1alpha1.datahub.events.Event) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.events.ListEventsRequest"></a>

### ListEventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| id | [string](#string) | repeated |  |
| cluster_id | [string](#string) | repeated |  |
| type | [EventType](#containersai.alameda.v1alpha1.datahub.events.EventType) | repeated |  |
| version | [EventVersion](#containersai.alameda.v1alpha1.datahub.events.EventVersion) | repeated |  |
| level | [EventLevel](#containersai.alameda.v1alpha1.datahub.events.EventLevel) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.events.ListEventsResponse"></a>

### ListEventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| events | [Event](#containersai.alameda.v1alpha1.datahub.events.Event) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/events/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/events/types.proto



<a name="containersai.alameda.v1alpha1.datahub.events.EventSource"></a>

### EventSource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |
| component | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.events.K8SObjectReference"></a>

### K8SObjectReference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |
| api_version | [string](#string) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.events.EventLevel"></a>

### EventLevel


| Name | Number | Description |
| ---- | ------ | ----------- |
| EVENT_LEVEL_UNDEFINED | 0 |  |
| EVENT_LEVEL_DEBUG | 1 |  |
| EVENT_LEVEL_INFO | 2 |  |
| EVENT_LEVEL_WARNING | 3 |  |
| EVENT_LEVEL_ERROR | 4 |  |
| EVENT_LEVEL_FATAL | 5 |  |



<a name="containersai.alameda.v1alpha1.datahub.events.EventType"></a>

### EventType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EVENT_TYPE_UNDEFINED | 0 |  |
| EVENT_TYPE_ALAMEDA_SCALER_CREATE | 1 |  |
| EVENT_TYPE_ALAMEDA_SCALER_DELETE | 2 |  |
| EVENT_TYPE_NODE_REGISTER | 3 |  |
| EVENT_TYPE_DEPLOYMENT_REGISTER | 4 |  |
| EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER | 5 |  |
| EVENT_TYPE_POD_REGISTER | 6 |  |
| EVENT_TYPE_NODE_DEREGISTER | 7 |  |
| EVENT_TYPE_DEPLOYMENT_DEREGISTER | 8 |  |
| EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER | 9 |  |
| EVENT_TYPE_POD_DEREGISTER | 10 |  |
| EVENT_TYPE_NODE_PREDICTION_CREATE | 11 |  |
| EVENT_TYPE_POD_PREDICTION_CREATE | 12 |  |
| EVENT_TYPE_VPA_RECOMMENDATION_CREATE | 13 |  |
| EVENT_TYPE_HPA_RECOMMENDATION_CREATE | 14 |  |
| EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE | 15 |  |
| EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE | 16 |  |
| EVENT_TYPE_ANOMALY_METRIC_DETECT | 17 |  |
| EVENT_TYPE_ANOMALY_ANALYSIS_CREATE | 18 |  |
| EVENT_TYPE_LICENSE | 19 |  |
| EVENT_TYPE_EMAIL_NOTIFICATION | 20 |  |
| EVENT_TYPE_ANOMALY_FORECAST_DETECT | 21 |  |
| EVENT_TYPE_ANOMALY_REALTIME_DETECT | 22 |  |



<a name="containersai.alameda.v1alpha1.datahub.events.EventVersion"></a>

### EventVersion


| Name | Number | Description |
| ---- | ------ | ----------- |
| EVENT_VERSION_UNDEFINED | 0 |  |
| EVENT_VERSION_V1 | 1 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/events/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/events/events.proto



<a name="containersai.alameda.v1alpha1.datahub.events.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| id | [string](#string) |  |  |
| cluster_id | [string](#string) |  |  |
| source | [EventSource](#containersai.alameda.v1alpha1.datahub.events.EventSource) |  |  |
| type | [EventType](#containersai.alameda.v1alpha1.datahub.events.EventType) |  |  |
| version | [EventVersion](#containersai.alameda.v1alpha1.datahub.events.EventVersion) |  |  |
| level | [EventLevel](#containersai.alameda.v1alpha1.datahub.events.EventLevel) |  |  |
| subject | [K8SObjectReference](#containersai.alameda.v1alpha1.datahub.events.K8SObjectReference) |  |  |
| message | [string](#string) |  |  |
| data | [string](#string) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/metrics/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/metrics/services.proto



<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest"></a>

### CreateApplicationMetricsRequest
Represents a request for creating metrics data of alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_metrics | [ApplicationMetric](#containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest"></a>

### CreateClusterMetricsRequest
Represents a request for creating metrics data of clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_metrics | [ClusterMetric](#containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest"></a>

### CreateControllerMetricsRequest
Represents a request for creating metrics data of controllers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| controller_metrics | [ControllerMetric](#containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest"></a>

### CreateMetricsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| metrics | [WriteMetric](#containersai.alameda.v1alpha1.datahub.metrics.WriteMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest"></a>

### CreateNamespaceMetricsRequest
Represents a request for creating metrics data of namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace_metrics | [NamespaceMetric](#containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest"></a>

### CreateNodeMetricsRequest
Represents a request for creating metrics data of nodes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_metrics | [NodeMetric](#containersai.alameda.v1alpha1.datahub.metrics.NodeMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest"></a>

### CreatePodMetricsRequest
Represents a request for creating metrics data of pods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_metrics | [PodMetric](#containersai.alameda.v1alpha1.datahub.metrics.PodMetric) | repeated |  |
| rate_range | [uint64](#uint64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest"></a>

### ListApplicationMetricsRequest
Represents a request for listing metric data of alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse"></a>

### ListApplicationMetricsResponse
Represents a response for a listing alameda scalers metric data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| application_metrics | [ApplicationMetric](#containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest"></a>

### ListClusterMetricsRequest
Represents a request for listing metric data of clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse"></a>

### ListClusterMetricsResponse
Represents a response for a listing clusters metrics request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| cluster_metrics | [ClusterMetric](#containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest"></a>

### ListControllerMetricsRequest
Represents a request for listing metric data of controllers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse"></a>

### ListControllerMetricsResponse
Represents a response for a listing controllers metric data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| controller_metrics | [ControllerMetric](#containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest"></a>

### ListMetricsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse"></a>

### ListMetricsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| metrics | [Metric](#containersai.alameda.v1alpha1.datahub.metrics.Metric) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest"></a>

### ListNamespaceMetricsRequest
Represents a request for listing metric data of namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse"></a>

### ListNamespaceMetricsResponse
Represents a response for a listing namespaces metric data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| namespace_metrics | [NamespaceMetric](#containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest"></a>

### ListNodeMetricsRequest
Represents a request for listing metric data of nodes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse"></a>

### ListNodeMetricsResponse
Represents a response for a listing nodes metrics request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_metrics | [NodeMetric](#containersai.alameda.v1alpha1.datahub.metrics.NodeMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest"></a>

### ListPodMetricsRequest
Represents a request for listing metric data of pods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| rate_range | [uint64](#uint64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse"></a>

### ListPodMetricsResponse
Represents a response for a listing pods metric data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_metrics | [PodMetric](#containersai.alameda.v1alpha1.datahub.metrics.PodMetric) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/metrics/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/metrics/types.proto



<a name="containersai.alameda.v1alpha1.datahub.metrics.Metric"></a>

### Metric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| metric_data | [MetricData](#containersai.alameda.v1alpha1.datahub.metrics.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.MetricData"></a>

### MetricData
Represents a piece of metreic data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| read_data | [containersai.alameda.v1alpha1.datahub.common.ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/metrics/metrics.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/metrics/metrics.proto



<a name="containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric"></a>

### ApplicationMetric
Represents metric data of a alameda scaler.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric"></a>

### ClusterMetric
Represents metric data of a cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric"></a>

### ContainerMetric
Represents metric data of a container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric"></a>

### ControllerMetric
Represents metric data of a controller.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric"></a>

### NamespaceMetric
Represents metric data of a namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.NodeMetric"></a>

### NodeMetric
Represents metric data of a node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.PodMetric"></a>

### PodMetric
Represents metric data of a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| container_metrics | [ContainerMetric](#containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.metrics.WriteMetric"></a>

### WriteMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| write_data | [containersai.alameda.v1alpha1.datahub.common.WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/applications/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/applications/services.proto



<a name="containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest"></a>

### CreateApplicationsRequest
Represents a request for adding alameda scalers that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| applications | [WriteApplication](#containersai.alameda.v1alpha1.datahub.applications.WriteApplication) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest"></a>

### DeleteApplicationsRequest
Represents a request for stopping predicting alameda sclaers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| applications | [DeleteApplication](#containersai.alameda.v1alpha1.datahub.applications.DeleteApplication) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest"></a>

### ListApplicationsRequest
Represents a request for listing alameda scalers that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| applications | [ReadApplication](#containersai.alameda.v1alpha1.datahub.applications.ReadApplication) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse"></a>

### ListApplicationsResponse
Represents a response for a listing alameda scalers request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| applications | [Application](#containersai.alameda.v1alpha1.datahub.applications.Application) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/applications/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/applications/types.proto



<a name="containersai.alameda.v1alpha1.datahub.applications.Application"></a>

### Application
Represents a dataset of private alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| application_data | [ApplicationData](#containersai.alameda.v1alpha1.datahub.applications.ApplicationData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.ApplicationData"></a>

### ApplicationData
Represents a private alameda scaler.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| read_data | [containersai.alameda.v1alpha1.datahub.common.ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/applications/applications.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/applications/applications.proto



<a name="containersai.alameda.v1alpha1.datahub.applications.DeleteApplication"></a>

### DeleteApplication
Represents the condition of deleting alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| where_condition | [containersai.alameda.v1alpha1.datahub.common.Condition](#containersai.alameda.v1alpha1.datahub.common.Condition) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.ReadApplication"></a>

### ReadApplication
Represents the condition of reading alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| where_condition | [containersai.alameda.v1alpha1.datahub.common.Condition](#containersai.alameda.v1alpha1.datahub.common.Condition) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.applications.WriteApplication"></a>

### WriteApplication
Represents the data of alameda scaler which is to be created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| write_data | [containersai.alameda.v1alpha1.datahub.common.WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/recommendations/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/recommendations/services.proto



<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest"></a>

### CreateApplicationRecommendationsRequest
Represents a request for creating alameda scalers&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_recommendations | [ApplicationRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ApplicationRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest"></a>

### CreateClusterRecommendationsRequest
Represents a request for creating clusters&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_recommendations | [ClusterRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ClusterRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest"></a>

### CreateControllerRecommendationsRequest
Represents a request for creating controllers&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| controller_recommendations | [ControllerRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest"></a>

### CreateNamespaceRecommendationsRequest
Represents a request for creating namespaces&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace_recommendations | [NamespaceRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NamespaceRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest"></a>

### CreateNodeRecommendationsRequest
Represents a request for creating nodes&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_recommendations | [NodeRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NodeRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest"></a>

### CreatePodRecommendationsRequest
Represents a request for creating pods&#39; recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_recommendations | [PodRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.PodRecommendation) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest"></a>

### CreateRecommendationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| recommendations | [WriteRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.WriteRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest"></a>

### ListApplicationRecommendationsRequest
Represents a request for listing recommendations of alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse"></a>

### ListApplicationRecommendationsResponse
Represents a response for listing alameda scalers recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| application_recommendations | [ApplicationRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ApplicationRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest"></a>

### ListClusterRecommendationsRequest
Represents a request for listing recommendations of clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse"></a>

### ListClusterRecommendationsResponse
Represents a response for listing clusters recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| cluster_recommendations | [ClusterRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ClusterRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest"></a>

### ListControllerRecommendationsRequest
Represents a request for listing recommendations of controllers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse"></a>

### ListControllerRecommendationsResponse
Represents a response for listing controllers recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| controller_recommendations | [ControllerRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest"></a>

### ListNamespaceRecommendationsRequest
Represents a request for listing recommendations of namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse"></a>

### ListNamespaceRecommendationsResponse
Represents a response for listing namespaces recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| namespace_recommendations | [NamespaceRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NamespaceRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest"></a>

### ListNodeRecommendationsRequest
Represents a request for listing recommendations of nodes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse"></a>

### ListNodeRecommendationsResponse
Represents a response for listing nodes recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_recommendations | [NodeRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.NodeRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest"></a>

### ListPodRecommendationsRequest
Represents a request for listing recommendations of pods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse"></a>

### ListPodRecommendationsResponse
Represents a response for listing pods recommendations request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_recommendations | [PodRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.PodRecommendation) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest"></a>

### ListRecommendationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse"></a>

### ListRecommendationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| recommendations | [Recommendation](#containersai.alameda.v1alpha1.datahub.recommendations.Recommendation) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/recommendations/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/recommendations/types.proto



<a name="containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec"></a>

### ControllerRecommendedSpec
Represents a private spec of a controller recommendation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_replicas | [int32](#int32) |  |  |
| desired_replicas | [int32](#int32) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| current_cpu_requests | [double](#double) |  |  |
| current_mem_requests | [double](#double) |  |  |
| current_cpu_limits | [double](#double) |  |  |
| current_mem_limits | [double](#double) |  |  |
| desired_cpu_limits | [double](#double) |  |  |
| desired_mem_limits | [double](#double) |  |  |
| total_cost | [double](#double) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s"></a>

### ControllerRecommendedSpecK8s



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_replicas | [int32](#int32) |  |  |
| desired_replicas | [int32](#int32) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.Recommendation"></a>

### Recommendation
Represents a set of container resource configuration recommendations of a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| recommendation_data | [RecommendationData](#containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.RecommendationData"></a>

### RecommendationData
Represents a piece of recommendation data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| read_data | [containersai.alameda.v1alpha1.datahub.common.ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType"></a>

### ControllerRecommendedType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CRT_UNDEFINED | 0 |  |
| PRIMITIVE | 1 |  |
| K8S | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/recommendations/recommendations.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/recommendations/recommendations.proto



<a name="containersai.alameda.v1alpha1.datahub.recommendations.ApplicationRecommendation"></a>

### ApplicationRecommendation
Represents resource configuration recommendations of a alameda scaler.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |
| recommended_spec | [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec) |  |  |
| recommended_spec_k8s | [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ClusterRecommendation"></a>

### ClusterRecommendation
Represents resource configuration recommendations of a cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |
| recommended_spec | [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec) |  |  |
| recommended_spec_k8s | [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ContainerRecommendation"></a>

### ContainerRecommendation
Represents a resource configuration recommendation made by the AI Engine.

It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial strage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| limit_recommendations | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| request_recommendations | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_limit_recommendations | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_request_recommendations | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendation"></a>

### ControllerRecommendation
Represents resource configuration recommendations of a controller.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |
| recommended_spec | [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec) |  |  |
| recommended_spec_k8s | [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.NamespaceRecommendation"></a>

### NamespaceRecommendation
Represents resource configuration recommendations of a namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |
| recommended_spec | [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec) |  |  |
| recommended_spec_k8s | [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.NodeRecommendation"></a>

### NodeRecommendation
Represents resource configuration recommendations of a node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| recommended_type | [ControllerRecommendedType](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType) |  |  |
| recommended_spec | [ControllerRecommendedSpec](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpec) |  |  |
| recommended_spec_k8s | [ControllerRecommendedSpecK8s](#containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedSpecK8s) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.PodRecommendation"></a>

### PodRecommendation
Represents a set of container resource configuration recommendations of a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| apply_recommendation_now | [bool](#bool) |  |  |
| assign_pod_policy | [containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy](#containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy) |  |  |
| container_recommendations | [ContainerRecommendation](#containersai.alameda.v1alpha1.datahub.recommendations.ContainerRecommendation) | repeated |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| top_controller | [containersai.alameda.v1alpha1.datahub.resources.Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) |  |  |
| recommendation_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.recommendations.WriteRecommendation"></a>

### WriteRecommendation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| write_data | [containersai.alameda.v1alpha1.datahub.common.WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/common/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/common/types.proto


 


<a name="containersai.alameda.v1alpha1.datahub.common.ColumnType"></a>

### ColumnType
Represents the field type of a record which queried from datahub.

| Name | Number | Description |
| ---- | ------ | ----------- |
| COLUMNTYPE_UDEFINED | 0 |  |
| COLUMNTYPE_TAG | 1 |  |
| COLUMNTYPE_FIELD | 2 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.DataType"></a>

### DataType
Represents the datahub specified data type.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DATATYPE_UNDEFINED | 0 |  |
| DATATYPE_BOOL | 1 |  |
| DATATYPE_INT | 2 |  |
| DATATYPE_INT8 | 3 |  |
| DATATYPE_INT16 | 4 |  |
| DATATYPE_INT32 | 5 |  |
| DATATYPE_INT64 | 6 |  |
| DATATYPE_UINT | 7 |  |
| DATATYPE_UINT8 | 8 |  |
| DATATYPE_UINT16 | 9 |  |
| DATATYPE_UINT32 | 10 |  |
| DATATYPE_UTIN64 | 11 |  |
| DATATYPE_FLOAT32 | 12 |  |
| DATATYPE_FLOAT64 | 13 |  |
| DATATYPE_STRING | 14 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.DatabaseType"></a>

### DatabaseType
Represents a specified database whcih is to query.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 |  |
| INFLUXDB | 1 |  |
| PROMETHEUS | 2 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.FunctionType"></a>

### FunctionType
Represents the functional query type of datahub.

| Name | Number | Description |
| ---- | ------ | ----------- |
| FUNCTIONTYPE_UNDEFINED | 0 |  |
| FUNCTIONTYPE_COUNT | 1 | Aggregation function

Returns the number of non-null field values. |
| FUNCTIONTYPE_DISTINCT | 2 | Returns the list of unique field values. |
| FUNCTIONTYPE_INTEGRAL | 3 | Returns the area under the curve for subsequent field values. |
| FUNCTIONTYPE_MEAN | 4 | Returns the arithmetic mean (average) of field values. |
| FUNCTIONTYPE_MEDIAN | 5 | Returns the middle value from a sorted list of field values. |
| FUNCTIONTYPE_MODE | 6 | Returns the most frequent value in a list of field values. |
| FUNCTIONTYPE_SPREAD | 7 | Returns the difference between the minimum and maximum field values. |
| FUNCTIONTYPE_STDDEV | 8 | Returns the standard deviation of field values. |
| FUNCTIONTYPE_SUM | 9 | Returns the sum of field values. |
| FUNCTIONTYPE_BOTTOM | 10 | Selector function

Returns the smallest N field values. |
| FUNCTIONTYPE_FIRST | 11 | Returns the field value with the oldest timestamp. |
| FUNCTIONTYPE_LAST | 12 | Returns the field value with the most recent timestamp. |
| FUNCTIONTYPE_MAX | 13 | Returns the greatest field value. |
| FUNCTIONTYPE_MIN | 14 | Returns the lowest field value. |
| FUNCTIONTYPE_PERCENTILE | 15 | Returns the Nth percentile field value. |
| FUNCTIONTYPE_SAMPLE | 16 | Returns a random sample of N field values. SAMPLE() uses reservoir sampling to generate the random points. |
| FUNCTIONTYPE_TOP | 17 | Returns the greatest N field values. |
| FUNCTIONTYPE_DERIVATIVE | 18 | Transformation function

Returns the rate of change between subsequent field values. |



<a name="containersai.alameda.v1alpha1.datahub.common.ResourceBoundary"></a>

### ResourceBoundary
Represents the amount resources that a Kubernete object is allowed to use.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_BOUNDARY_UNDEFINED | 0 |  |
| RESOURCE_RAW | 1 |  |
| RESOURCE_UPPER_BOUND | 2 |  |
| RESOURCE_LOWER_BOUND | 3 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.ResourceQuota"></a>

### ResourceQuota
Represents the constraints that limit aggretage resource consumption per Kubernete object.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_QUOTA_UNDEFINED | 0 |  |
| RESOURCE_LIMIT | 1 |  |
| RESOURCE_REQUEST | 2 |  |
| RESOURCE_INITIAL_LIMIT | 3 |  |
| RESOURCE_INITIAL_REQUEST | 4 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/common/metrics.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/common/metrics.proto



<a name="containersai.alameda.v1alpha1.datahub.common.MetricData"></a>

### MetricData
Represents a piece of metreic data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| data | [Sample](#containersai.alameda.v1alpha1.datahub.common.Sample) | repeated | Data can be time series or non-time series. |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.common.Sample"></a>

### Sample
Represents a data point of time-series metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| num_value | [string](#string) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.common.MetricType"></a>

### MetricType
Metric type. A metric may be CPU, memory and etc.

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRICS_TYPE_UNDEFINED | 0 |  |
| CPU_SECONDS_TOTAL | 1 |  |
| CPU_CORES_ALLOCATABLE | 2 |  |
| CPU_MILLICORES_TOTAL | 3 |  |
| CPU_MILLICORES_AVAIL | 4 |  |
| CPU_MILLICORES_USAGE | 5 |  |
| CPU_MILLICORES_USAGE_PCT | 6 |  |
| MEMORY_BYTES_ALLOCATABLE | 7 |  |
| MEMORY_BYTES_TOTAL | 8 |  |
| MEMORY_BYTES_AVAIL | 9 |  |
| MEMORY_BYTES_USAGE | 10 |  |
| MEMORY_BYTES_USAGE_PCT | 11 |  |
| FS_BYTES_TOTAL | 12 |  |
| FS_BYTES_AVAIL | 13 |  |
| FS_BYTES_USAGE | 14 |  |
| FS_BYTES_USAGE_PCT | 15 |  |
| HTTP_REQUESTS_COUNT | 16 |  |
| HTTP_REQUESTS_TOTAL | 17 |  |
| HTTP_RESPONSE_COUNT | 18 |  |
| HTTP_RESPONSE_TOTAL | 19 |  |
| DISK_IO_SECONDS_TOTAL | 20 |  |
| DISK_IO_UTILIZATION | 21 |  |
| RESTARTS_TOTAL | 22 |  |
| UNSCHEDULABLE | 23 |  |
| HEALTH | 24 |  |
| POWER_USAGE_WATTS | 25 |  |
| TEMPERATURE_CELSIUS | 26 |  |
| DUTY_CYCLE | 27 |  |
| CURRENT_OFFSET | 28 |  |
| LAG | 29 |  |
| LATENCY | 30 |  |
| NUMBER | 31 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.ResourceName"></a>

### ResourceName
Represents Kubernetes resources which will be allocated to pods.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_NAME_UNDEFINED | 0 |  |
| CPU | 1 |  |
| MEMORY | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/common/queries.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/common/queries.proto



<a name="containersai.alameda.v1alpha1.datahub.common.Condition"></a>

### Condition
Represents a query condition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keys | [string](#string) | repeated |  |
| values | [string](#string) | repeated |  |
| operators | [string](#string) | repeated |  |
| types | [DataType](#containersai.alameda.v1alpha1.datahub.common.DataType) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.common.Function"></a>

### Function
Represents a datahub functional query, includes aggregation, selector and transformation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [FunctionType](#containersai.alameda.v1alpha1.datahub.common.FunctionType) |  |  |
| fields | [string](#string) | repeated |  |
| tags | [string](#string) | repeated |  |
| target | [string](#string) |  |  |
| regular_expression | [string](#string) |  |  |
| unit | [string](#string) |  |  |
| number | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.common.Into"></a>

### Into
Represents a query to a user-specified measurement.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| retention_policy | [string](#string) |  |  |
| measurement | [string](#string) |  |  |
| is_default_retention_policy | [bool](#bool) |  |  |
| is_all_measurements | [bool](#bool) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.common.QueryCondition"></a>

### QueryCondition
Represents a datahub query request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time_range | [TimeRange](#containersai.alameda.v1alpha1.datahub.common.TimeRange) |  |  |
| order | [QueryCondition.Order](#containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order) |  |  |
| function | [Function](#containersai.alameda.v1alpha1.datahub.common.Function) |  |  |
| into | [Into](#containersai.alameda.v1alpha1.datahub.common.Into) |  |  |
| where_clause | [string](#string) |  |  |
| where_condition | [Condition](#containersai.alameda.v1alpha1.datahub.common.Condition) | repeated |  |
| selects | [string](#string) | repeated |  |
| groups | [string](#string) | repeated |  |
| limit | [uint64](#uint64) |  |  |
| sub_query | [QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.common.TimeRange"></a>

### TimeRange
Represents a time range definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| timeout | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| step | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| aggregateFunction | [TimeRange.AggregateFunction](#containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction) |  |  |
| apply_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order"></a>

### QueryCondition.Order


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ASC | 1 |  |
| DESC | 2 |  |



<a name="containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction"></a>

### TimeRange.AggregateFunction


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| MAX | 1 |  |
| AVG | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/common/rawdata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/common/rawdata.proto



<a name="containersai.alameda.v1alpha1.datahub.common.Group"></a>

### Group
Represents a dataset which are collected that have the same attributes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| columns | [string](#string) | repeated |  |
| rows | [Row](#containersai.alameda.v1alpha1.datahub.common.Row) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.common.ReadData"></a>

### ReadData
Represents a dataset whcih is read from datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [Group](#containersai.alameda.v1alpha1.datahub.common.Group) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.common.Row"></a>

### Row
Represents a record of data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| values | [string](#string) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.common.WriteData"></a>

### WriteData
Represents a dataset which will be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| columns | [string](#string) | repeated |  |
| rows | [Row](#containersai.alameda.v1alpha1.datahub.common.Row) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/weavescope/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/weavescope/services.proto



<a name="containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest"></a>

### ListWeaveScopeContainersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| container_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest"></a>

### ListWeaveScopeHostsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest"></a>

### ListWeaveScopePodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse"></a>

### WeaveScopeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| rawdata | [string](#string) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/rawdata/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/rawdata/services.proto



<a name="containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest"></a>

### ReadRawdataRequest
Represents a request for reading rawdata from database.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_type | [containersai.alameda.v1alpha1.datahub.common.DatabaseType](#containersai.alameda.v1alpha1.datahub.common.DatabaseType) |  |  |
| queries | [Query](#containersai.alameda.v1alpha1.datahub.rawdata.Query) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse"></a>

### ReadRawdataResponse
Represents a response for listing rawdata from database.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| rawdata | [ReadRawdata](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest"></a>

### WriteRawdataRequest
Represents a request for writing rawdata to database.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_type | [containersai.alameda.v1alpha1.datahub.common.DatabaseType](#containersai.alameda.v1alpha1.datahub.common.DatabaseType) |  |  |
| rawdata | [WriteRawdata](#containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/rawdata/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/rawdata/types.proto



<a name="containersai.alameda.v1alpha1.datahub.rawdata.Query"></a>

### Query
Represents a general datahub query request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| table | [string](#string) |  |  |
| expression | [string](#string) |  |  |
| condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/rawdata/rawdata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/rawdata/rawdata.proto



<a name="containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata"></a>

### ReadRawdata
Represents a rawdata whcih is read from datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [Query](#containersai.alameda.v1alpha1.datahub.rawdata.Query) |  |  |
| columns | [string](#string) | repeated |  |
| groups | [containersai.alameda.v1alpha1.datahub.common.Group](#containersai.alameda.v1alpha1.datahub.common.Group) | repeated |  |
| rawdata | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata"></a>

### WriteRawdata
Represents a rawdata which will be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [string](#string) |  |  |
| table | [string](#string) |  |  |
| columns | [string](#string) | repeated |  |
| rows | [containersai.alameda.v1alpha1.datahub.common.Row](#containersai.alameda.v1alpha1.datahub.common.Row) | repeated |  |
| column_types | [containersai.alameda.v1alpha1.datahub.common.ColumnType](#containersai.alameda.v1alpha1.datahub.common.ColumnType) | repeated |  |
| data_types | [containersai.alameda.v1alpha1.datahub.common.DataType](#containersai.alameda.v1alpha1.datahub.common.DataType) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/licenses/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/licenses/services.proto



<a name="containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse"></a>

### GetLicenseResponse
Represents a response for reading a license information request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| license | [License](#containersai.alameda.v1alpha1.datahub.licenses.License) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/licenses/licenses.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/licenses/licenses.proto



<a name="containersai.alameda.v1alpha1.datahub.licenses.License"></a>

### License
Represents the information of a license.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Valid | [bool](#bool) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/scores/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/scores/services.proto



<a name="containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest"></a>

### CreateSimulatedSchedulingScoresRequest
Represents a request for adding scheduling scores produced by ai engine


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scores | [SimulatedSchedulingScore](#containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScore) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest"></a>

### ListSimulatedSchedulingScoresRequest
Represents a request for listing system scores of pod scheduled on node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse"></a>

### ListSimulatedSchedulingScoresResponse
Represents a response for listing system scores request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| scores | [SimulatedSchedulingScore](#containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScore) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/scores/scores.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/scores/scores.proto



<a name="containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScore"></a>

### SimulatedSchedulingScore
Represents a system score before and after pod scheduled on node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| score_before | [double](#double) |  |  |
| score_after | [double](#double) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/data/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/data/services.proto



<a name="containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest"></a>

### DeleteDataRequest
Represents a request for deleting data in datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| delete_data | [DeleteData](#containersai.alameda.v1alpha1.datahub.data.DeleteData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.data.ReadDataRequest"></a>

### ReadDataRequest
Represents a request for reading data from datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| read_data | [ReadData](#containersai.alameda.v1alpha1.datahub.data.ReadData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.data.ReadDataResponse"></a>

### ReadDataResponse
Represents a response for a reading data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| data | [Data](#containersai.alameda.v1alpha1.datahub.data.Data) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.data.WriteDataRequest"></a>

### WriteDataRequest
Represents a request for writing data to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| write_data | [WriteData](#containersai.alameda.v1alpha1.datahub.data.WriteData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.data.WriteMetaRequest"></a>

### WriteMetaRequest
Represents a request for writing data(none time-series) to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| write_meta | [WriteMeta](#containersai.alameda.v1alpha1.datahub.data.WriteMeta) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/data/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/data/types.proto



<a name="containersai.alameda.v1alpha1.datahub.data.Data"></a>

### Data
Represents a dataset of rawdata which will be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| rawdata | [Rawdata](#containersai.alameda.v1alpha1.datahub.data.Rawdata) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.data.Rawdata"></a>

### Rawdata
Represents a private alameda specified rawdata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| groups | [containersai.alameda.v1alpha1.datahub.common.Group](#containersai.alameda.v1alpha1.datahub.common.Group) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/data/data.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/data/data.proto



<a name="containersai.alameda.v1alpha1.datahub.data.DeleteData"></a>

### DeleteData
Represents the condition of deleting data in datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.data.ReadData"></a>

### ReadData
Represents the condition of reading data from datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.data.WriteData"></a>

### WriteData
Represents the data which is to be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| columns | [string](#string) | repeated |  |
| rows | [containersai.alameda.v1alpha1.datahub.common.Row](#containersai.alameda.v1alpha1.datahub.common.Row) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.data.WriteMeta"></a>

### WriteMeta
Represents the data(none time-series) which is to be written to datahub.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| measurement | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| condition | [containersai.alameda.v1alpha1.datahub.common.Condition](#containersai.alameda.v1alpha1.datahub.common.Condition) |  |  |
| columns | [string](#string) | repeated |  |
| rows | [containersai.alameda.v1alpha1.datahub.common.Row](#containersai.alameda.v1alpha1.datahub.common.Row) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/services.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest"></a>

### CreateApplicationsRequest
Represents a request for creating alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| applications | [Application](#containersai.alameda.v1alpha1.datahub.resources.Application) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest"></a>

### CreateClustersRequest
Represents a request for adding clusters that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusters | [Cluster](#containersai.alameda.v1alpha1.datahub.resources.Cluster) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest"></a>

### CreateControllersRequest
Represents a request for creating controllers to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| controllers | [Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest"></a>

### CreateNamespacesRequest
Represents a request for creating namespaces to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaces | [Namespace](#containersai.alameda.v1alpha1.datahub.resources.Namespace) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest"></a>

### CreateNodesRequest
Represents a request for adding nodes that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#containersai.alameda.v1alpha1.datahub.resources.Node) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest"></a>

### CreatePodsRequest
Represents a request for creating pods to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pods | [Pod](#containersai.alameda.v1alpha1.datahub.resources.Pod) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest"></a>

### DeleteApplicationsRequest
Represents a request for deleting alameda scalers data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest"></a>

### DeleteClustersRequest
Represents a request for stopping predicting clusters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest"></a>

### DeleteControllersRequest
Represents a request for stopping predicting controllers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest"></a>

### DeleteNamespacesRequest
Represents a request for stopping predicting namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest"></a>

### DeleteNodesRequest
Represents a request for stopping predicting nodes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest"></a>

### DeletePodsRequest
Represents a request for stopping predicting pods.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest"></a>

### ListApplicationsRequest
Represents a request for listing alameda scalers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse"></a>

### ListApplicationsResponse
Represents a response for a listing alameda scalers request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| applications | [Application](#containersai.alameda.v1alpha1.datahub.resources.Application) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest"></a>

### ListClustersRequest
Represents a request for listing clusters that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse"></a>

### ListClustersResponse
Represents a response for a listing clusters request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| clusters | [Cluster](#containersai.alameda.v1alpha1.datahub.resources.Cluster) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest"></a>

### ListControllersRequest
Represents a request for listing controllers that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse"></a>

### ListControllersResponse
Represents a response for a listing controllers request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| controllers | [Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest"></a>

### ListNamespacesRequest
Represents a request for listing namespaces that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse"></a>

### ListNamespacesResponse
Represents a response for a listing namespaces request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| namespaces | [Namespace](#containersai.alameda.v1alpha1.datahub.resources.Namespace) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest"></a>

### ListNodesRequest
Represents a request for listing nodes that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse"></a>

### ListNodesResponse
Represents a response for a listing nodes request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| nodes | [Node](#containersai.alameda.v1alpha1.datahub.resources.Node) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest"></a>

### ListPodsRequest
Represents a request for listing pods that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| scaling_tool | [ScalingTool](#containersai.alameda.v1alpha1.datahub.resources.ScalingTool) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse"></a>

### ListPodsResponse
Represents a response for a listing pods request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pods | [Pod](#containersai.alameda.v1alpha1.datahub.resources.Pod) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/metadata.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.ObjectMeta"></a>

### ObjectMeta
Represents the private metadata of alameda object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| node_name | [string](#string) |  |  |
| cluster_name | [string](#string) |  |  |
| uid | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.OwnerReference"></a>

### OwnerReference
Represents the owner of Kubernetes object. The owned objects are called dependents of the owner object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.resources.Kind"></a>

### Kind
Represents Kubernetes resource kind.

| Name | Number | Description |
| ---- | ------ | ----------- |
| KIND_UNDEFINED | 0 |  |
| DEPLOYMENT | 1 |  |
| DEPLOYMENTCONFIG | 2 |  |
| STATEFULSET | 3 |  |
| ALAMEDASCALER | 4 |  |



<a name="containersai.alameda.v1alpha1.datahub.resources.ScalingTool"></a>

### ScalingTool
Represents the scaling tool for managing Kubernetes resources.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SCALING_TOOL_UNDEFINED | 0 |  |
| NONE | 1 |  |
| VPA | 2 |  |
| HPA | 3 |  |
| CA | 4 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/policies.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/policies.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy"></a>

### AssignPodPolicy
Represents a recommended pod-to-node assignment. (i.e. pod placement)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| node_priority | [NodePriority](#containersai.alameda.v1alpha1.datahub.resources.NodePriority) |  |  |
| node_selector | [Selector](#containersai.alameda.v1alpha1.datahub.resources.Selector) |  |  |
| node_name | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.NodePriority"></a>

### NodePriority
Represents the priority of a node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [string](#string) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Selector"></a>

### Selector
Represents a Kubernetes label selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [Selector.SelectorEntry](#containersai.alameda.v1alpha1.datahub.resources.Selector.SelectorEntry) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Selector.SelectorEntry"></a>

### Selector.SelectorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.resources.RecommendationPolicy"></a>

### RecommendationPolicy
Recommendation policy. A policy may be either stable or compact.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RECOMMENDATION_POLICY_UNDEFINED | 0 |  |
| STABLE | 1 |  |
| COMPACT | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/types.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.AlamedaApplicationSpec"></a>

### AlamedaApplicationSpec
Represents the private alameda applcation specification.






<a name="containersai.alameda.v1alpha1.datahub.resources.AlamedaControllerSpec"></a>

### AlamedaControllerSpec
Represents the private alameda controller specification.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alameda_scaler | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| scaling_tool | [ScalingTool](#containersai.alameda.v1alpha1.datahub.resources.ScalingTool) |  |  |
| policy | [RecommendationPolicy](#containersai.alameda.v1alpha1.datahub.resources.RecommendationPolicy) |  |  |
| min_replicas | [int32](#int32) |  |  |
| max_replicas | [int32](#int32) |  |  |
| enable_recommendation_execution | [bool](#bool) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.AlamedaNodeSpec"></a>

### AlamedaNodeSpec
Represents the private alameda node specification.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#containersai.alameda.v1alpha1.datahub.resources.Provider) |  |  |
| machineset_name | [string](#string) |  |  |
| machineset_namespace | [string](#string) |  |  |
| role_master | [bool](#bool) |  |  |
| role_worker | [bool](#bool) |  |  |
| role_infra | [bool](#bool) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.AlamedaPodSpec"></a>

### AlamedaPodSpec
Represents the private alameda pod specification.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alameda_scaler | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| scaling_tool | [ScalingTool](#containersai.alameda.v1alpha1.datahub.resources.ScalingTool) |  |  |
| used_recommendation_id | [string](#string) |  |  |
| policy | [RecommendationPolicy](#containersai.alameda.v1alpha1.datahub.resources.RecommendationPolicy) |  |  |
| alameda_scaler_resources | [ResourceRequirements](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Capacity"></a>

### Capacity
Represents the capacity of a Kubernetes node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu_cores | [int64](#int64) |  |  |
| memory_bytes | [int64](#int64) |  |  |
| network_megabits_per_second | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Provider"></a>

### Provider
The information of cloud service provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |
| instance_type | [string](#string) |  |  |
| region | [string](#string) |  |  |
| zone | [string](#string) |  |  |
| os | [string](#string) |  |  |
| role | [string](#string) |  |  |
| instance_id | [string](#string) |  |  |
| storage_size | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements"></a>

### ResourceRequirements
ResourceRequirements describes the compute resource requirements


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limits | [ResourceRequirements.LimitsEntry](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.LimitsEntry) | repeated | limits describes the maximum amount of compute resources allowed use enum ResourceName as key of the map which defined in common |
| requests | [ResourceRequirements.RequestsEntry](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.RequestsEntry) | repeated | requests describes the minimum amount of compute resources required use enum ResourceName as key of the map which defined in common |






<a name="containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.LimitsEntry"></a>

### ResourceRequirements.LimitsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements.RequestsEntry"></a>

### ResourceRequirements.RequestsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/status.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.ContainerState"></a>

### ContainerState
ContainerState holds a possible state of container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| waiting | [ContainerStateWaiting](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateWaiting) |  |  |
| running | [ContainerStateRunning](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateRunning) |  |  |
| terminated | [ContainerStateTerminated](#containersai.alameda.v1alpha1.datahub.resources.ContainerStateTerminated) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ContainerStateRunning"></a>

### ContainerStateRunning
ContainerStateRunning is a running state of a container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| started_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ContainerStateTerminated"></a>

### ContainerStateTerminated
ContainerStateTerminated is a terminated state of a container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exit_code | [int32](#int32) |  |  |
| reason | [string](#string) |  |  |
| message | [string](#string) |  |  |
| started_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| finished_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ContainerStateWaiting"></a>

### ContainerStateWaiting
ContainerStateWaiting is a waiting state of a container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reason | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.ContainerStatus"></a>

### ContainerStatus
ContainerStatus contains details for the current status of this container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [ContainerState](#containersai.alameda.v1alpha1.datahub.resources.ContainerState) |  |  |
| last_termination_state | [ContainerState](#containersai.alameda.v1alpha1.datahub.resources.ContainerState) |  |  |
| restart_count | [int32](#int32) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.PodStatus"></a>

### PodStatus
PodStatus represents information about the status of a pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phase | [PodPhase](#containersai.alameda.v1alpha1.datahub.resources.PodPhase) |  |  |
| message | [string](#string) |  |  |
| reason | [string](#string) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.resources.PodPhase"></a>

### PodPhase
The valid statuses of pods.

| Name | Number | Description |
| ---- | ------ | ----------- |
| POD_PHASE_UNDEFINED | 0 |  |
| PENDING | 1 |  |
| RUNNING | 2 |  |
| SUCCEEDED | 3 |  |
| FAILED | 4 |  |
| UNKNOWN | 5 |  |
| COMPLETED | 6 |  |
| CRASHLOOPBACKOFF | 7 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/resources/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/resources/resources.proto



<a name="containersai.alameda.v1alpha1.datahub.resources.Application"></a>

### Application
Represents a alameda scaler.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| alameda_application_spec | [AlamedaApplicationSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaApplicationSpec) |  |  |
| controllers | [Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Cluster"></a>

### Cluster
Represents a Kubernetes cluster.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Container"></a>

### Container
Represents a container and its containing limit and requeset configurations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| resources | [ResourceRequirements](#containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements) |  |  |
| status | [ContainerStatus](#containersai.alameda.v1alpha1.datahub.resources.ContainerStatus) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Controller"></a>

### Controller
Represents a Kubernetes namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| replicas | [int32](#int32) |  |  |
| spec_replicas | [int32](#int32) |  |  |
| alameda_controller_spec | [AlamedaControllerSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaControllerSpec) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Namespace"></a>

### Namespace
Represents a Kubernetes namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Node"></a>

### Node
Represents a Kubernetes node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| machine_start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| capacity | [Capacity](#containersai.alameda.v1alpha1.datahub.resources.Capacity) |  |  |
| alameda_node_spec | [AlamedaNodeSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaNodeSpec) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.resources.Pod"></a>

### Pod
Represents a Kubernetes pod.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| resource_link | [string](#string) |  |  |
| app_name | [string](#string) |  |  |
| app_part_of | [string](#string) |  |  |
| alameda_pod_spec | [AlamedaPodSpec](#containersai.alameda.v1alpha1.datahub.resources.AlamedaPodSpec) |  |  |
| top_controller | [Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) |  |  |
| status | [PodStatus](#containersai.alameda.v1alpha1.datahub.resources.PodStatus) |  |  |
| containers | [Container](#containersai.alameda.v1alpha1.datahub.resources.Container) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/keycodes/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/keycodes/services.proto



<a name="containersai.alameda.v1alpha1.datahub.keycodes.ActivateRegistrationDataRequest"></a>

### ActivateRegistrationDataRequest
Represents a request for activating license signature data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [string](#string) |  | example: &#34;FaS3fivdY7d1xEYxmSa&#43;mg==&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeRequest"></a>

### AddKeycodeRequest
Represents a request for adding a keycode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keycode | [string](#string) |  | example: &#34;A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeResponse"></a>

### AddKeycodeResponse
Represents a response for adding a keycode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| keycode | [Keycode](#containersai.alameda.v1alpha1.datahub.keycodes.Keycode) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.DeleteKeycodeRequest"></a>

### DeleteKeycodeRequest
Represents a request for deleting a keycode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keycode | [string](#string) |  | example: &#34;A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.GenerateRegistrationDataResponse"></a>

### GenerateRegistrationDataResponse
Represents a request for generating license registration data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| data | [string](#string) |  | example: &#34;FaS3fivdY7d1xEYxmSa&#43;mg==&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesRequest"></a>

### ListKeycodesRequest
Represents a request for retrieving keycodes detailed information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keycodes | [string](#string) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesResponse"></a>

### ListKeycodesResponse
Represents a response for a retrieving keycodes detailed information request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| keycodes | [Keycode](#containersai.alameda.v1alpha1.datahub.keycodes.Keycode) | repeated |  |
| summary | [Keycode](#containersai.alameda.v1alpha1.datahub.keycodes.Keycode) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/keycodes/keycodes.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/keycodes/keycodes.proto



<a name="containersai.alameda.v1alpha1.datahub.keycodes.Keycode"></a>

### Keycode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keycode | [string](#string) |  | example: &#34;A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST&#34;` |
| keycode_type | [string](#string) |  | example: &#34;Regular/Trial&#34;` |
| keycode_version | [int32](#int32) |  | example: &#34;2&#34;` |
| apply_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example: &#34;2018-01-01T00:00:00Z&#34;` |
| expire_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example: &#34;2018-12-31T11:59:59Z&#34;` |
| license_state | [string](#string) |  | example: &#34;Valid/Invalid/Expired&#34;` |
| registered | [bool](#bool) |  | example: &#34;false&#34;` |
| capacity | [Capacity](#containersai.alameda.v1alpha1.datahub.keycodes.Capacity) |  |  |
| functionality | [Functionality](#containersai.alameda.v1alpha1.datahub.keycodes.Functionality) |  |  |
| retention | [Retention](#containersai.alameda.v1alpha1.datahub.keycodes.Retention) |  |  |
| service_agreement | [ServiceAgreement](#containersai.alameda.v1alpha1.datahub.keycodes.ServiceAgreement) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/keycodes/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/keycodes/types.proto



<a name="containersai.alameda.v1alpha1.datahub.keycodes.Capacity"></a>

### Capacity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [int32](#int32) |  | example: &#34;-1&#34; |
| hosts | [int32](#int32) |  | example: &#34;20&#34; |
| disks | [int32](#int32) |  | example: &#34;200&#34; |
| cpus | [int32](#int32) |  | example: &#34;64&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.Functionality"></a>

### Functionality



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disk_prophet | [bool](#bool) |  | example: &#34;true&#34; |
| workload | [bool](#bool) |  | example: &#34;true&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.Retention"></a>

### Retention



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid_month | [int32](#int32) |  | example: &#34;0&#34; |
| years | [int32](#int32) |  | example: &#34;1&#34; |






<a name="containersai.alameda.v1alpha1.datahub.keycodes.ServiceAgreement"></a>

### ServiceAgreement






 

 

 

 



<a name="alameda_api/v1alpha1/datahub/server.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/server.proto


 

 

 


<a name="containersai.alameda.v1alpha1.datahub.DatahubService"></a>

### DatahubService
Service for providing data stored in the backend

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateApps | [applications.CreateApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ListApps | [applications.ListApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest) | [applications.ListApplicationsResponse](#containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse) |  |
| DeleteApps | [applications.DeleteApplicationsRequest](#containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ReadData | [data.ReadDataRequest](#containersai.alameda.v1alpha1.datahub.data.ReadDataRequest) | [data.ReadDataResponse](#containersai.alameda.v1alpha1.datahub.data.ReadDataResponse) | Used to read data based on alameda specific schemas |
| WriteData | [data.WriteDataRequest](#containersai.alameda.v1alpha1.datahub.data.WriteDataRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to write data based on alameda specific schemas |
| DeleteData | [data.DeleteDataRequest](#containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to delete data based on alameda specific schemas |
| WriteMeta | [data.WriteMetaRequest](#containersai.alameda.v1alpha1.datahub.data.WriteMetaRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to write metadata based on alameda specific schemas |
| CreateEvents | [events.CreateEventsRequest](#containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create alameda specific events |
| ListEvents | [events.ListEventsRequest](#containersai.alameda.v1alpha1.datahub.events.ListEventsRequest) | [events.ListEventsResponse](#containersai.alameda.v1alpha1.datahub.events.ListEventsResponse) | Used to list alameda specific events |
| CreateGpuPredictions | [gpu.CreateGpuPredictionsRequest](#containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create GPU predictions |
| ListGpus | [gpu.ListGpusRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest) | [gpu.ListGpusResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse) | Used to list GPU need to be predicted |
| ListGpuMetrics | [gpu.ListGpuMetricsRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest) | [gpu.ListGpuMetricsResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse) | Used to list GPU metrics data |
| ListGpuPredictions | [gpu.ListGpuPredictionsRequest](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest) | [gpu.ListGpuPredictionsResponse](#containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse) | Used to list GPU predictions |
| AddKeycode | [keycodes.AddKeycodeRequest](#containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeRequest) | [keycodes.AddKeycodeResponse](#containersai.alameda.v1alpha1.datahub.keycodes.AddKeycodeResponse) | Used to add a keycode |
| ListKeycodes | [keycodes.ListKeycodesRequest](#containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesRequest) | [keycodes.ListKeycodesResponse](#containersai.alameda.v1alpha1.datahub.keycodes.ListKeycodesResponse) | Used to retrieve keycodes detailed information |
| DeleteKeycode | [keycodes.DeleteKeycodeRequest](#containersai.alameda.v1alpha1.datahub.keycodes.DeleteKeycodeRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to delete a keycode |
| GenerateRegistrationData | [.google.protobuf.Empty](#google.protobuf.Empty) | [keycodes.GenerateRegistrationDataResponse](#containersai.alameda.v1alpha1.datahub.keycodes.GenerateRegistrationDataResponse) | Used to generate license registration data |
| ActivateRegistrationData | [keycodes.ActivateRegistrationDataRequest](#containersai.alameda.v1alpha1.datahub.keycodes.ActivateRegistrationDataRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to activate license signature data |
| GetLicense | [.google.protobuf.Empty](#google.protobuf.Empty) | [licenses.GetLicenseResponse](#containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse) | Used to get datahub license information |
| CreateMetrics | [metrics.CreateMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreatePodMetrics | [metrics.CreatePodMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateControllerMetrics | [metrics.CreateControllerMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateApplicationMetrics | [metrics.CreateApplicationMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateNamespaceMetrics | [metrics.CreateNamespaceMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateNodeMetrics | [metrics.CreateNodeMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateClusterMetrics | [metrics.CreateClusterMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ListMetrics | [metrics.ListMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest) | [metrics.ListMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse) |  |
| ListPodMetrics | [metrics.ListPodMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest) | [metrics.ListPodMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse) | Used to list pod metric data |
| ListControllerMetrics | [metrics.ListControllerMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest) | [metrics.ListControllerMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse) | Used to list controller metric data |
| ListApplicationMetrics | [metrics.ListApplicationMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest) | [metrics.ListApplicationMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse) | Used to list alameda scaler metric data |
| ListNamespaceMetrics | [metrics.ListNamespaceMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest) | [metrics.ListNamespaceMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse) | Used to list namespace metric data |
| ListNodeMetrics | [metrics.ListNodeMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest) | [metrics.ListNodeMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse) | Used to list node metric data |
| ListClusterMetrics | [metrics.ListClusterMetricsRequest](#containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest) | [metrics.ListClusterMetricsResponse](#containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse) | Used to list cluster metric data |
| Ping | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.rpc.Status](#google.rpc.Status) | Used to check if datahub is still alive |
| CreatePlannings | [plannings.CreatePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreatePodPlannings | [plannings.CreatePodPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateControllerPlannings | [plannings.CreateControllerPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateApplicationPlannings | [plannings.CreateApplicationPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateNamespacePlannings | [plannings.CreateNamespacePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateNodePlannings | [plannings.CreateNodePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateClusterPlannings | [plannings.CreateClusterPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ListPlannings | [plannings.ListPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest) | [plannings.ListPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse) |  |
| ListPodPlannings | [plannings.ListPodPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest) | [plannings.ListPodPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse) |  |
| ListControllerPlannings | [plannings.ListControllerPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest) | [plannings.ListControllerPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse) |  |
| ListApplicationPlannings | [plannings.ListApplicationPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest) | [plannings.ListApplicationPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse) |  |
| ListNamespacePlannings | [plannings.ListNamespacePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest) | [plannings.ListNamespacePlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse) |  |
| ListNodePlannings | [plannings.ListNodePlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest) | [plannings.ListNodePlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse) |  |
| ListClusterPlannings | [plannings.ListClusterPlanningsRequest](#containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest) | [plannings.ListClusterPlanningsResponse](#containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse) |  |
| CreatePredictions | [predictions.CreatePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreatePodPredictions | [predictions.CreatePodPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of pods |
| CreateControllerPredictions | [predictions.CreateControllerPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of controllers |
| CreateApplicationPredictions | [predictions.CreateApplicationPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of alameda scalers |
| CreateNamespacePredictions | [predictions.CreateNamespacePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of namespaces |
| CreateNodePredictions | [predictions.CreateNodePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of nodes |
| CreateClusterPredictions | [predictions.CreateClusterPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of clusters |
| ListPredictions | [predictions.ListPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest) | [predictions.ListPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse) |  |
| ListPodPredictions | [predictions.ListPodPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest) | [predictions.ListPodPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse) | Used to list pod predictions |
| ListControllerPredictions | [predictions.ListControllerPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest) | [predictions.ListControllerPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse) | Used to list controller predictions |
| ListApplicationPredictions | [predictions.ListApplicationPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest) | [predictions.ListApplicationPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse) | Used to list alameda scaler predictions |
| ListNamespacePredictions | [predictions.ListNamespacePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest) | [predictions.ListNamespacePredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse) | Used to list namespace predictions |
| ListNodePredictions | [predictions.ListNodePredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest) | [predictions.ListNodePredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse) | Used to list node predictions |
| ListClusterPredictions | [predictions.ListClusterPredictionsRequest](#containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest) | [predictions.ListClusterPredictionsResponse](#containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse) | Used to list cluster predictions |
| ReadRawdata | [rawdata.ReadRawdataRequest](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest) | [rawdata.ReadRawdataResponse](#containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse) |  |
| WriteRawdata | [rawdata.WriteRawdataRequest](#containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateRecommendations | [recommendations.CreateRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreatePodRecommendations | [recommendations.CreatePodRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of pods |
| CreateControllerRecommendations | [recommendations.CreateControllerRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of controllers |
| CreateApplicationRecommendations | [recommendations.CreateApplicationRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of alameda scalers |
| CreateNamespaceRecommendations | [recommendations.CreateNamespaceRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of namespaces |
| CreateNodeRecommendations | [recommendations.CreateNodeRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of nodes |
| CreateClusterRecommendations | [recommendations.CreateClusterRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of clusters |
| ListRecommendations | [recommendations.ListRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest) | [recommendations.ListRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse) |  |
| ListPodRecommendations | [recommendations.ListPodRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest) | [recommendations.ListPodRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse) | Used to list pod recommenations |
| ListAvailablePodRecommendations | [recommendations.ListPodRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest) | [recommendations.ListPodRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse) | Used to list available pod recommenations |
| ListControllerRecommendations | [recommendations.ListControllerRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest) | [recommendations.ListControllerRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse) | Used to list controller recommenations |
| ListApplicationRecommendations | [recommendations.ListApplicationRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest) | [recommendations.ListApplicationRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse) | Used to list alameda scaler recommenations |
| ListNamespaceRecommendations | [recommendations.ListNamespaceRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest) | [recommendations.ListNamespaceRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse) | Used to list namespace recommenations |
| ListNodeRecommendations | [recommendations.ListNodeRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest) | [recommendations.ListNodeRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse) | Used to list node recommenations |
| ListClusterRecommendations | [recommendations.ListClusterRecommendationsRequest](#containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest) | [recommendations.ListClusterRecommendationsResponse](#containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse) | Used to list cluster recommenations |
| CreatePods | [resources.CreatePodsRequest](#containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add pods that need to be predicted |
| CreateControllers | [resources.CreateControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add controllers that need to be predicted |
| CreateApplications | [resources.CreateApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add alameda scalers that need to be predicted |
| CreateNamespaces | [resources.CreateNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add namespaces that need to be predicted |
| CreateNodes | [resources.CreateNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add nodes that need to be predicted |
| CreateClusters | [resources.CreateClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add clusters that need to be predicted |
| ListPods | [resources.ListPodsRequest](#containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest) | [resources.ListPodsResponse](#containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse) | Used to list pods need to be predicted |
| ListControllers | [resources.ListControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest) | [resources.ListControllersResponse](#containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse) | Used to list contollers need to be predicted |
| ListApplications | [resources.ListApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest) | [resources.ListApplicationsResponse](#containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse) | Used to list alameda scalers need to be predicted |
| ListNamespaces | [resources.ListNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest) | [resources.ListNamespacesResponse](#containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse) | Used to list namespaces need to be predicted |
| ListNodes | [resources.ListNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest) | [resources.ListNodesResponse](#containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse) | Used to list nodes&#39; information |
| ListClusters | [resources.ListClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest) | [resources.ListClustersResponse](#containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse) | Used to list clusters&#39; information |
| DeletePods | [resources.DeletePodsRequest](#containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the pods |
| DeleteControllers | [resources.DeleteControllersRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the controllers |
| DeleteApplications | [resources.DeleteApplicationsRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the applications |
| DeleteNamespaces | [resources.DeleteNamespacesRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the namespaces |
| DeleteNodes | [resources.DeleteNodesRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the nodes |
| DeleteClusters | [resources.DeleteClustersRequest](#containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to stop generating predictions for the clusters |
| CreateSchemas | [schemas.CreateSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ListSchemas | [schemas.ListSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest) | [schemas.ListSchemasResponse](#containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse) |  |
| DeleteSchemas | [schemas.DeleteSchemasRequest](#containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| CreateSimulatedSchedulingScores | [scores.CreateSimulatedSchedulingScoresRequest](#containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest) | [.google.rpc.Status](#google.rpc.Status) |  |
| ListSimulatedSchedulingScores | [scores.ListSimulatedSchedulingScoresRequest](#containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest) | [scores.ListSimulatedSchedulingScoresResponse](#containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse) | Used to list system scores |
| ListWeaveScopeHosts | [weavescope.ListWeaveScopeHostsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| GetWeaveScopeHostDetails | [weavescope.ListWeaveScopeHostsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| ListWeaveScopePods | [weavescope.ListWeaveScopePodsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| GetWeaveScopePodDetails | [weavescope.ListWeaveScopePodsRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| ListWeaveScopeContainers | [weavescope.ListWeaveScopeContainersRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| ListWeaveScopeContainersByHostname | [weavescope.ListWeaveScopeContainersRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| ListWeaveScopeContainersByImage | [weavescope.ListWeaveScopeContainersRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |
| GetWeaveScopeContainerDetails | [weavescope.ListWeaveScopeContainersRequest](#containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest) | [weavescope.WeaveScopeResponse](#containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse) |  |

 



<a name="alameda_api/v1alpha1/datahub/plannings/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/plannings/services.proto



<a name="containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest"></a>

### CreateApplicationPlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_plannings | [ApplicationPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest"></a>

### CreateClusterPlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_plannings | [ClusterPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest"></a>

### CreateControllerPlanningsRequest
Represents a request for creating a controller&#39;s plannings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| controller_plannings | [ControllerPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest"></a>

### CreateNamespacePlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace_plannings | [NamespacePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest"></a>

### CreateNodePlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_plannings | [NodePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NodePlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest"></a>

### CreatePlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| plannings | [WritePlanning](#containersai.alameda.v1alpha1.datahub.plannings.WritePlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest"></a>

### CreatePodPlanningsRequest
Represents a request for creating a pod&#39;s plannings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_plannings | [PodPlanning](#containersai.alameda.v1alpha1.datahub.plannings.PodPlanning) | repeated |  |
| granularity | [int64](#int64) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest"></a>

### ListApplicationPlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse"></a>

### ListApplicationPlanningsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| application_plannings | [ApplicationPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest"></a>

### ListClusterPlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse"></a>

### ListClusterPlanningsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| cluster_plannings | [ClusterPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest"></a>

### ListControllerPlanningsRequest
Represents a request for listing plannings of controllers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse"></a>

### ListControllerPlanningsResponse
Represents a response for listing controller plannings request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| controller_plannings | [ControllerPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest"></a>

### ListNamespacePlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse"></a>

### ListNamespacePlanningsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| namespace_plannings | [NamespacePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest"></a>

### ListNodePlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse"></a>

### ListNodePlanningsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_plannings | [NodePlanning](#containersai.alameda.v1alpha1.datahub.plannings.NodePlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest"></a>

### ListPlanningsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse"></a>

### ListPlanningsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| plannings | [RawPlanning](#containersai.alameda.v1alpha1.datahub.plannings.RawPlanning) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest"></a>

### ListPodPlanningsRequest
Represents a request for listing plannings of pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) | repeated |  |
| granularity | [int64](#int64) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse"></a>

### ListPodPlanningsResponse
Represents a response for listing pod plannings request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_plannings | [PodPlanning](#containersai.alameda.v1alpha1.datahub.plannings.PodPlanning) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/plannings/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/plannings/types.proto



<a name="containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningSpec"></a>

### ControllerPlanningSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_replicas | [int32](#int32) |  |  |
| desired_replicas | [int32](#int32) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| current_cpu_requests | [double](#double) |  |  |
| current_mem_requests | [double](#double) |  |  |
| current_cpu_limits | [double](#double) |  |  |
| current_mem_limits | [double](#double) |  |  |
| desired_cpu_limits | [double](#double) |  |  |
| desired_mem_limits | [double](#double) |  |  |
| total_cost | [double](#double) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningSpecK8s"></a>

### ControllerPlanningSpecK8s



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_replicas | [int32](#int32) |  |  |
| desired_replicas | [int32](#int32) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.PlanningData"></a>

### PlanningData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| read_data | [containersai.alameda.v1alpha1.datahub.common.ReadData](#containersai.alameda.v1alpha1.datahub.common.ReadData) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.RawPlanning"></a>

### RawPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| planning_data | [PlanningData](#containersai.alameda.v1alpha1.datahub.plannings.PlanningData) | repeated |  |





 


<a name="containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningType"></a>

### ControllerPlanningType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CPT_UNDEFINED | 0 |  |
| CPT_PRIMITIVE | 1 |  |
| CPT_K8S | 2 |  |



<a name="containersai.alameda.v1alpha1.datahub.plannings.PlanningType"></a>

### PlanningType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PT_UNDEFINED | 0 |  |
| PT_RECOMMENDATION | 1 |  |
| PT_PLANNING | 2 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/plannings/plannings.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/plannings/plannings.proto



<a name="containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning"></a>

### ApplicationPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plannings | [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning"></a>

### ClusterPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plannings | [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning"></a>

### ContainerPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| limit_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| request_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_limit_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_request_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning"></a>

### ControllerPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| kind | [containersai.alameda.v1alpha1.datahub.resources.Kind](#containersai.alameda.v1alpha1.datahub.resources.Kind) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plannings | [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning"></a>

### NamespacePlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plannings | [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.NodePlanning"></a>

### NodePlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plannings | [Planning](#containersai.alameda.v1alpha1.datahub.plannings.Planning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.Planning"></a>

### Planning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| request_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_limit_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |
| initial_request_plannings | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.PodPlanning"></a>

### PodPlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_meta | [containersai.alameda.v1alpha1.datahub.resources.ObjectMeta](#containersai.alameda.v1alpha1.datahub.resources.ObjectMeta) |  |  |
| planning_type | [PlanningType](#containersai.alameda.v1alpha1.datahub.plannings.PlanningType) |  |  |
| planning_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |
| total_cost | [double](#double) |  |  |
| apply_planning_now | [bool](#bool) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| assign_pod_policy | [containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy](#containersai.alameda.v1alpha1.datahub.resources.AssignPodPolicy) |  |  |
| top_controller | [containersai.alameda.v1alpha1.datahub.resources.Controller](#containersai.alameda.v1alpha1.datahub.resources.Controller) |  |  |
| container_plannings | [ContainerPlanning](#containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.plannings.WritePlanning"></a>

### WritePlanning



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| write_data | [containersai.alameda.v1alpha1.datahub.common.WriteData](#containersai.alameda.v1alpha1.datahub.common.WriteData) |  |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/gpu/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/gpu/services.proto



<a name="containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest"></a>

### CreateGpuPredictionsRequest
Represents a request for creating predictions of graphics processing units&#39; metric data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gpu_predictions | [GpuPrediction](#containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest"></a>

### ListGpuMetricsRequest
Represents a request for listing metric data of graphics processing units.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| metric_types | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) | repeated |  |
| host | [string](#string) |  |  |
| minor_number | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse"></a>

### ListGpuMetricsResponse
Represents a response for a listing graphics processing units metric data request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| gpu_metrics | [GpuMetric](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetric) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest"></a>

### ListGpuPredictionsRequest
Represents a list of predicted metric data of graphics processing units.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| host | [string](#string) |  |  |
| minor_number | [string](#string) |  |  |
| granularity | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| prediction_id | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse"></a>

### ListGpuPredictionsResponse
Represents a response for a listing predictions of graphics processing units request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| gpu_predictions | [GpuPrediction](#containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest"></a>

### ListGpusRequest
Represents a request for listing graphics processing units that need to be predicted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [containersai.alameda.v1alpha1.datahub.common.QueryCondition](#containersai.alameda.v1alpha1.datahub.common.QueryCondition) |  |  |
| host | [string](#string) |  |  |
| minor_number | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse"></a>

### ListGpusResponse
Represents a response for a listing graphics processing units request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| gpus | [Gpu](#containersai.alameda.v1alpha1.datahub.gpu.Gpu) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/gpu/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/gpu/types.proto



<a name="containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata"></a>

### GpuMetadata
Represents the metadata of a graphics processing unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |
| instance | [string](#string) |  |  |
| job | [string](#string) |  |  |
| minor_number | [string](#string) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.GpuSpec"></a>

### GpuSpec
Represents the spec of a graphics processing unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memory_total | [float](#float) |  | Total memory in bytes |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/gpu/gpu.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/gpu/gpu.proto



<a name="containersai.alameda.v1alpha1.datahub.gpu.Gpu"></a>

### Gpu
Represents a graphics processing unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| uuid | [string](#string) |  |  |
| metadata | [GpuMetadata](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata) |  |  |
| spec | [GpuSpec](#containersai.alameda.v1alpha1.datahub.gpu.GpuSpec) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.GpuMetric"></a>

### GpuMetric
Represents metric data of a graphics processing unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| uuid | [string](#string) |  |  |
| metadata | [GpuMetadata](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata) |  |  |
| metric_data | [containersai.alameda.v1alpha1.datahub.common.MetricData](#containersai.alameda.v1alpha1.datahub.common.MetricData) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction"></a>

### GpuPrediction
Represents a list of predicted metrics data of a graphics processing unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| uuid | [string](#string) |  |  |
| metadata | [GpuMetadata](#containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata) |  |  |
| predicted_raw_data | [containersai.alameda.v1alpha1.datahub.predictions.MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_upperbound_data | [containersai.alameda.v1alpha1.datahub.predictions.MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |
| predicted_lowerbound_data | [containersai.alameda.v1alpha1.datahub.predictions.MetricData](#containersai.alameda.v1alpha1.datahub.predictions.MetricData) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/schemas/services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/schemas/services.proto



<a name="containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest"></a>

### CreateSchemasRequest
Represents a request for creating datahub schemas.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schemas | [Schema](#containersai.alameda.v1alpha1.datahub.schemas.Schema) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest"></a>

### DeleteSchemasRequest
Represents a request for deleting datahub schemas.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest"></a>

### ListSchemasRequest
Represents a request for listing datahub schemas.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse"></a>

### ListSchemasResponse
Represents a response for a listing datahub schemas request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| schemas | [Schema](#containersai.alameda.v1alpha1.datahub.schemas.Schema) | repeated |  |





 

 

 

 



<a name="alameda_api/v1alpha1/datahub/schemas/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/schemas/types.proto



<a name="containersai.alameda.v1alpha1.datahub.schemas.Column"></a>

### Column
Represents a data record.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| required | [bool](#bool) |  |  |
| column_type | [containersai.alameda.v1alpha1.datahub.common.ColumnType](#containersai.alameda.v1alpha1.datahub.common.ColumnType) |  |  |
| data_type | [containersai.alameda.v1alpha1.datahub.common.DataType](#containersai.alameda.v1alpha1.datahub.common.DataType) |  |  |






<a name="containersai.alameda.v1alpha1.datahub.schemas.Measurement"></a>

### Measurement
Represents the information of measurment which datahub will write data in InfluxDB.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| metric_type | [containersai.alameda.v1alpha1.datahub.common.MetricType](#containersai.alameda.v1alpha1.datahub.common.MetricType) |  |  |
| resource_boundary | [containersai.alameda.v1alpha1.datahub.common.ResourceBoundary](#containersai.alameda.v1alpha1.datahub.common.ResourceBoundary) |  |  |
| resource_quota | [containersai.alameda.v1alpha1.datahub.common.ResourceQuota](#containersai.alameda.v1alpha1.datahub.common.ResourceQuota) |  |  |
| is_ts | [bool](#bool) |  |  |
| columns | [Column](#containersai.alameda.v1alpha1.datahub.schemas.Column) | repeated |  |






<a name="containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta"></a>

### SchemaMeta
Represents the private metadata of datahub schema.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scope | [Scope](#containersai.alameda.v1alpha1.datahub.schemas.Scope) |  |  |
| category | [string](#string) |  |  |
| type | [string](#string) |  |  |





 


<a name="containersai.alameda.v1alpha1.datahub.schemas.Scope"></a>

### Scope


| Name | Number | Description |
| ---- | ------ | ----------- |
| SCOPE_UNDEFINED | 0 |  |
| SCOPE_APPLICATION | 1 |  |
| SCOPE_CONFIG | 2 |  |
| SCOPE_EXECUTION | 3 |  |
| SCOPE_FEDEMETER | 4 |  |
| SCOPE_METRIC | 5 |  |
| SCOPE_PLANNING | 6 |  |
| SCOPE_PREDICTION | 7 |  |
| SCOPE_RECOMMENDATION | 8 |  |
| SCOPE_RESOURCE | 9 |  |
| SCOPE_TARGET | 10 |  |


 

 

 



<a name="alameda_api/v1alpha1/datahub/schemas/schemas.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alameda_api/v1alpha1/datahub/schemas/schemas.proto



<a name="containersai.alameda.v1alpha1.datahub.schemas.Schema"></a>

### Schema
Represents the private datahub specific schema of data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_meta | [SchemaMeta](#containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta) |  |  |
| measurements | [Measurement](#containersai.alameda.v1alpha1.datahub.schemas.Measurement) | repeated |  |





 

 

 

 



<a name="datahub/recommendation/v1alpha2/recommendation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/recommendation/v1alpha2/recommendation.proto
This file has messages related to recommendations of containers, pods, and nodes


<a name="containersai.datahub.recommendation.v1alpha2.AssignPodPolicy"></a>

### AssignPodPolicy
Represents a recommended pod-to-node assignment (i.e. pod placement)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| node_priority | [containersai.datahub.resource.pod.assign.v1alpha2.NodePriority](#containersai.datahub.resource.pod.assign.v1alpha2.NodePriority) |  |  |
| node_selector | [containersai.datahub.resource.pod.assign.v1alpha2.Selector](#containersai.datahub.resource.pod.assign.v1alpha2.Selector) |  |  |
| node_name | [string](#string) |  |  |






<a name="containersai.datahub.recommendation.v1alpha2.ContainerRecommendation"></a>

### ContainerRecommendation
Represents a resource configuration recommendation

It includes recommended limits and requests for the initial stage (a container which is just started) and after the initial stage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| limit_recommendations | [ContainerRecommendation.LimitRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.LimitRecommendationsEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |
| request_recommendations | [ContainerRecommendation.RequestRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.RequestRecommendationsEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |
| initial_limit_recommendations | [ContainerRecommendation.InitialLimitRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialLimitRecommendationsEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |
| initial_request_recommendations | [ContainerRecommendation.InitialRequestRecommendationsEntry](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialRequestRecommendationsEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |






<a name="containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialLimitRecommendationsEntry"></a>

### ContainerRecommendation.InitialLimitRecommendationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.InitialRequestRecommendationsEntry"></a>

### ContainerRecommendation.InitialRequestRecommendationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.LimitRecommendationsEntry"></a>

### ContainerRecommendation.LimitRecommendationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.recommendation.v1alpha2.ContainerRecommendation.RequestRecommendationsEntry"></a>

### ContainerRecommendation.RequestRecommendationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.recommendation.v1alpha2.PodRecommendation"></a>

### PodRecommendation
Represents a set of container resource configuration recommenations of a pod


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| apply_recommendation_now | [bool](#bool) |  |  |
| assign_pod_policy | [AssignPodPolicy](#containersai.datahub.recommendation.v1alpha2.AssignPodPolicy) |  |  |
| container_recommendations | [ContainerRecommendation](#containersai.datahub.recommendation.v1alpha2.ContainerRecommendation) | repeated |  |





 


<a name="containersai.datahub.recommendation.v1alpha2.RecommendationPolicy"></a>

### RecommendationPolicy
Recommendation policy. A policy may be either stable or compact.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RECOMMENDATIONPOLICY_UNDEFINED | 0 |  |
| STABLE | 1 |  |
| COMPACT | 2 |  |


 

 

 



<a name="datahub/score/v1alpha2/score.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/score/v1alpha2/score.proto
This file has messages and services related to Containers.ai


<a name="containersai.datahub.score.v1alpha2.SimulatedSchedulingScore"></a>

### SimulatedSchedulingScore
Represents a system score before and after pod scheduled on node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| score_before | [double](#double) |  |  |
| score_after | [double](#double) |  |  |





 

 

 

 



<a name="datahub/v1alpha2/datahub.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/v1alpha2/datahub.proto
This file has messages related to Kubernetes metadata


<a name="containersai.datahub.v1alpha2.CreateNodePredictionsRequest"></a>

### CreateNodePredictionsRequest
Represents a request for creating predictions of a node metric data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_predictions | [containersai.datahub.prediction.v1alpha2.NodePrediction](#containersai.datahub.prediction.v1alpha2.NodePrediction) | repeated |  |






<a name="containersai.datahub.v1alpha2.CreateNodesRequest"></a>

### CreateNodesRequest
Represents a request for adding nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [containersai.datahub.resource.v1alpha2.Node](#containersai.datahub.resource.v1alpha2.Node) | repeated |  |






<a name="containersai.datahub.v1alpha2.CreatePodPredictionsRequest"></a>

### CreatePodPredictionsRequest
Represents a request for creating predictions of containers&#39; metric data belonging to a pod


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_predictions | [containersai.datahub.prediction.v1alpha2.PodPrediction](#containersai.datahub.prediction.v1alpha2.PodPrediction) | repeated |  |






<a name="containersai.datahub.v1alpha2.CreatePodRecommendationsRequest"></a>

### CreatePodRecommendationsRequest
Represents a request for creating a pod&#39;s recommendation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pod_recommendations | [containersai.datahub.recommendation.v1alpha2.PodRecommendation](#containersai.datahub.recommendation.v1alpha2.PodRecommendation) | repeated |  |






<a name="containersai.datahub.v1alpha2.CreatePodsRequest"></a>

### CreatePodsRequest
Represents a request for creating pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pods | [containersai.datahub.resource.v1alpha2.Pod](#containersai.datahub.resource.v1alpha2.Pod) | repeated |  |






<a name="containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest"></a>

### CreateSimulatedSchedulingScoresRequest
Represents a request for adding scheduling scores produced by ai engine


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scores | [containersai.datahub.score.v1alpha2.SimulatedSchedulingScore](#containersai.datahub.score.v1alpha2.SimulatedSchedulingScore) | repeated |  |






<a name="containersai.datahub.v1alpha2.DeleteNodesRequest"></a>

### DeleteNodesRequest
Represents a request for stoping nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [containersai.datahub.resource.v1alpha2.Node](#containersai.datahub.resource.v1alpha2.Node) | repeated |  |






<a name="containersai.datahub.v1alpha2.DeletePodsRequest"></a>

### DeletePodsRequest
Represents a request for deleting pods data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pods | [containersai.datahub.resource.v1alpha2.Pod](#containersai.datahub.resource.v1alpha2.Pod) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListNodeMetricsRequest"></a>

### ListNodeMetricsRequest
Represents a request for listing metrics of nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_names | [string](#string) | repeated |  |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListNodeMetricsResponse"></a>

### ListNodeMetricsResponse
Represents a response for a listing metrics of nodes request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_metrics | [containersai.datahub.metric.v1alpha2.NodeMetric](#containersai.datahub.metric.v1alpha2.NodeMetric) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListNodePredictionsRequest"></a>

### ListNodePredictionsRequest
Represents a request for listing predictions of nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_names | [string](#string) | repeated |  |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListNodePredictionsResponse"></a>

### ListNodePredictionsResponse
Represents a response for a listing predictions of nodes request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| node_predictions | [containersai.datahub.prediction.v1alpha2.NodePrediction](#containersai.datahub.prediction.v1alpha2.NodePrediction) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListNodesRequest"></a>

### ListNodesRequest
Represents a request for listing nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_predicted | [bool](#bool) |  |  |






<a name="containersai.datahub.v1alpha2.ListNodesResponse"></a>

### ListNodesResponse
Represents a response for a listing nodes request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| nodes | [containersai.datahub.resource.v1alpha2.Node](#containersai.datahub.resource.v1alpha2.Node) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListPodMetricsRequest"></a>

### ListPodMetricsRequest
Represents a request for listing metric data of pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListPodMetricsResponse"></a>

### ListPodMetricsResponse
Represents a response for a listing metrics of pods request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_metrics | [containersai.datahub.metric.v1alpha2.PodMetric](#containersai.datahub.metric.v1alpha2.PodMetric) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListPodPredictionsRequest"></a>

### ListPodPredictionsRequest
Represents a request for listing predictions of pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListPodPredictionsResponse"></a>

### ListPodPredictionsResponse
Represents a response for a listing predictions of pods request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_predictions | [containersai.datahub.prediction.v1alpha2.PodPrediction](#containersai.datahub.prediction.v1alpha2.PodPrediction) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListPodRecommendationsRequest"></a>

### ListPodRecommendationsRequest
Represents a request for listing recommendations of pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListPodRecommendationsResponse"></a>

### ListPodRecommendationsResponse
Represents a response for listing pod recommendations request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pod_recommendations | [containersai.datahub.recommendation.v1alpha2.PodRecommendation](#containersai.datahub.recommendation.v1alpha2.PodRecommendation) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListPodsRequest"></a>

### ListPodsRequest
Represents a request for listing pods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scaler | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| is_predicted | [bool](#bool) |  |  |






<a name="containersai.datahub.v1alpha2.ListPodsResponse"></a>

### ListPodsResponse
Represents a response for a listing pods request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| pods | [containersai.datahub.resource.v1alpha2.Pod](#containersai.datahub.resource.v1alpha2.Pod) | repeated |  |






<a name="containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest"></a>

### ListSimulatedSchedulingScoresRequest
Represents a request for listing system scores of pod scheduled on node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_condition | [QueryCondition](#containersai.datahub.v1alpha2.QueryCondition) |  |  |






<a name="containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse"></a>

### ListSimulatedSchedulingScoresResponse
Represents a response for listing system scores request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  |  |
| scores | [containersai.datahub.score.v1alpha2.SimulatedSchedulingScore](#containersai.datahub.score.v1alpha2.SimulatedSchedulingScore) | repeated |  |






<a name="containersai.datahub.v1alpha2.QueryCondition"></a>

### QueryCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time_range | [containersai.datahub.metric.v1alpha2.TimeRange](#containersai.datahub.metric.v1alpha2.TimeRange) |  |  |
| order | [QueryCondition.Order](#containersai.datahub.v1alpha2.QueryCondition.Order) |  |  |
| limit | [uint64](#uint64) |  |  |






<a name="containersai.datahub.v1alpha2.UpdateNodesRequest"></a>

### UpdateNodesRequest
Represents a request for upating nodes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated_nodes | [UpdateNodesRequest.UpdatedNode](#containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode) | repeated |  |






<a name="containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode"></a>

### UpdateNodesRequest.UpdatedNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is required |
| is_predicted | [UpdateNodesRequest.UpdatedNode.IsPredictedWrap](#containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode.IsPredictedWrap) |  |  |






<a name="containersai.datahub.v1alpha2.UpdateNodesRequest.UpdatedNode.IsPredictedWrap"></a>

### UpdateNodesRequest.UpdatedNode.IsPredictedWrap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_predicted | [bool](#bool) |  |  |






<a name="containersai.datahub.v1alpha2.UpdatePodsRequest"></a>

### UpdatePodsRequest
Represents a request for updating pods data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated_pods | [UpdatePodsRequest.UpdatedPod](#containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod) | repeated |  |






<a name="containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod"></a>

### UpdatePodsRequest.UpdatedPod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  | namespace and name in containersai.datahub.resource.metadata.v1alpha2.NamespacedName are required |
| is_predicted | [UpdatePodsRequest.UpdatedPod.IsPredictedWrap](#containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod.IsPredictedWrap) |  |  |






<a name="containersai.datahub.v1alpha2.UpdatePodsRequest.UpdatedPod.IsPredictedWrap"></a>

### UpdatePodsRequest.UpdatedPod.IsPredictedWrap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_predicted | [bool](#bool) |  |  |





 


<a name="containersai.datahub.v1alpha2.QueryCondition.Order"></a>

### QueryCondition.Order


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASC | 0 |  |
| DESC | 1 |  |


 

 


<a name="containersai.datahub.v1alpha2.DatahubService"></a>

### DatahubService
Service for providing data stored in the backend

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListPodMetrics | [ListPodMetricsRequest](#containersai.datahub.v1alpha2.ListPodMetricsRequest) | [ListPodMetricsResponse](#containersai.datahub.v1alpha2.ListPodMetricsResponse) | Used to list pod metric data |
| ListNodeMetrics | [ListNodeMetricsRequest](#containersai.datahub.v1alpha2.ListNodeMetricsRequest) | [ListNodeMetricsResponse](#containersai.datahub.v1alpha2.ListNodeMetricsResponse) | Used to list node metric data |
| ListPods | [ListPodsRequest](#containersai.datahub.v1alpha2.ListPodsRequest) | [ListPodsResponse](#containersai.datahub.v1alpha2.ListPodsResponse) | Used to list pods |
| ListNodes | [ListNodesRequest](#containersai.datahub.v1alpha2.ListNodesRequest) | [ListNodesResponse](#containersai.datahub.v1alpha2.ListNodesResponse) | Used to list nodes |
| ListPodPredictions | [ListPodPredictionsRequest](#containersai.datahub.v1alpha2.ListPodPredictionsRequest) | [ListPodPredictionsResponse](#containersai.datahub.v1alpha2.ListPodPredictionsResponse) | Used to list pod predictions |
| ListNodePredictions | [ListNodePredictionsRequest](#containersai.datahub.v1alpha2.ListNodePredictionsRequest) | [ListNodePredictionsResponse](#containersai.datahub.v1alpha2.ListNodePredictionsResponse) | Used to list node predictions |
| ListPodRecommendations | [ListPodRecommendationsRequest](#containersai.datahub.v1alpha2.ListPodRecommendationsRequest) | [ListPodRecommendationsResponse](#containersai.datahub.v1alpha2.ListPodRecommendationsResponse) | Used to list pod recommenations |
| ListSimulatedSchedulingScores | [ListSimulatedSchedulingScoresRequest](#containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest) | [ListSimulatedSchedulingScoresResponse](#containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse) | Used to list system scores |
| CreatePods | [CreatePodsRequest](#containersai.datahub.v1alpha2.CreatePodsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add pods that need to be predicted |
| CreateNodes | [CreateNodesRequest](#containersai.datahub.v1alpha2.CreateNodesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to add nodes |
| CreatePodPredictions | [CreatePodPredictionsRequest](#containersai.datahub.v1alpha2.CreatePodPredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of pods |
| CreateNodePredictions | [CreateNodePredictionsRequest](#containersai.datahub.v1alpha2.CreateNodePredictionsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create predictions of nodes |
| CreatePodRecommendations | [CreatePodRecommendationsRequest](#containersai.datahub.v1alpha2.CreatePodRecommendationsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create recommendations of pods |
| CreateSimulatedSchedulingScores | [CreateSimulatedSchedulingScoresRequest](#containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to create scores of system |
| UpdatePods | [UpdatePodsRequest](#containersai.datahub.v1alpha2.UpdatePodsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to update info of pods |
| UpdateNodes | [UpdateNodesRequest](#containersai.datahub.v1alpha2.UpdateNodesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to update info of nodes |
| DeletePods | [DeletePodsRequest](#containersai.datahub.v1alpha2.DeletePodsRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to delete info of pods |
| DeleteNodes | [DeleteNodesRequest](#containersai.datahub.v1alpha2.DeleteNodesRequest) | [.google.rpc.Status](#google.rpc.Status) | Used to delete info of nodes |

 



<a name="datahub/metric/v1alpha2/metric.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/metric/v1alpha2/metric.proto
This file has messages related to metric data of containers, pods, and nodes


<a name="containersai.datahub.metric.v1alpha2.ContainerMetric"></a>

### ContainerMetric
Represents metric data of a container


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| metric_data | [ContainerMetric.MetricDataEntry](#containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry) | repeated | use MetricType as key |






<a name="containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry"></a>

### ContainerMetric.MetricDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.metric.v1alpha2.MetricData"></a>

### MetricData
Represents a piece of metreic data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Sample](#containersai.datahub.metric.v1alpha2.Sample) | repeated | data can be time series or non-time series |






<a name="containersai.datahub.metric.v1alpha2.NodeMetric"></a>

### NodeMetric
Represents metric data of a node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| metric_data | [NodeMetric.MetricDataEntry](#containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry) | repeated | use MetricType as key |






<a name="containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry"></a>

### NodeMetric.MetricDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.metric.v1alpha2.PodMetric"></a>

### PodMetric
Represents metric data of a pod


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| container_metrics | [ContainerMetric](#containersai.datahub.metric.v1alpha2.ContainerMetric) | repeated |  |






<a name="containersai.datahub.metric.v1alpha2.Sample"></a>

### Sample
Represents a data point of time-series metric data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| num_value | [string](#string) |  |  |






<a name="containersai.datahub.metric.v1alpha2.TimeRange"></a>

### TimeRange
Represents a time range definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| step | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |





 


<a name="containersai.datahub.metric.v1alpha2.MetricType"></a>

### MetricType
Metric type. A metric may be either CPU or memory.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 |  |
| CPU_USAGE_SECONDS_PERCENTAGE | 1 |  |
| MEMORY_USAGE_BYTES | 2 |  |


 

 

 



<a name="datahub/resource/pod/assign/v1alpha2/assign.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/resource/pod/assign/v1alpha2/assign.proto
This file has messages related to nodes


<a name="containersai.datahub.resource.pod.assign.v1alpha2.NodePriority"></a>

### NodePriority
Represents the priority of a node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [string](#string) | repeated |  |






<a name="containersai.datahub.resource.pod.assign.v1alpha2.Selector"></a>

### Selector
Represents a Kubernetes label selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [Selector.SelectorEntry](#containersai.datahub.resource.pod.assign.v1alpha2.Selector.SelectorEntry) | repeated |  |






<a name="containersai.datahub.resource.pod.assign.v1alpha2.Selector.SelectorEntry"></a>

### Selector.SelectorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="datahub/resource/v1alpha2/resource.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/resource/v1alpha2/resource.proto
This file has messages related to compute resources


<a name="containersai.datahub.resource.v1alpha2.Container"></a>

### Container
Represents a container and its containing limit and requeset configurations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| limit_resource | [Container.LimitResourceEntry](#containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |
| request_resource | [Container.RequestResourceEntry](#containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |






<a name="containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry"></a>

### Container.LimitResourceEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry"></a>

### Container.RequestResourceEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.resource.v1alpha2.Node"></a>

### Node
Represents a Kubernetes node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| is_predicted | [bool](#bool) |  |  |






<a name="containersai.datahub.resource.v1alpha2.Pod"></a>

### Pod
Represents a Kubernetes pod


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| resource_link | [string](#string) |  |  |
| containers | [Container](#containersai.datahub.resource.v1alpha2.Container) | repeated |  |
| is_predicted | [bool](#bool) |  |  |
| scaler | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| node_name | [string](#string) |  |  |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| policy | [containersai.datahub.recommendation.v1alpha2.RecommendationPolicy](#containersai.datahub.recommendation.v1alpha2.RecommendationPolicy) |  |  |





 

 

 

 



<a name="datahub/resource/metadata/v1alpha2/metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/resource/metadata/v1alpha2/metadata.proto
This file has messages related to predictions of containers, pods, and nodes


<a name="containersai.datahub.resource.metadata.v1alpha2.NamespacedName"></a>

### NamespacedName
Represents kubernetes resource with namespace and name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 

 



<a name="datahub/prediction/v1alpha2/prediction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## datahub/prediction/v1alpha2/prediction.proto
This file has messages related to predictions of containers, pods, and nodes


<a name="containersai.datahub.prediction.v1alpha2.ContainerPrediction"></a>

### ContainerPrediction
Represents a list of predicted metric data of a container


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| predicted_raw_data | [ContainerPrediction.PredictedRawDataEntry](#containersai.datahub.prediction.v1alpha2.ContainerPrediction.PredictedRawDataEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |






<a name="containersai.datahub.prediction.v1alpha2.ContainerPrediction.PredictedRawDataEntry"></a>

### ContainerPrediction.PredictedRawDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.prediction.v1alpha2.NodePrediction"></a>

### NodePrediction
Represents a list of predicted metric data of a node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| predicted_raw_data | [NodePrediction.PredictedRawDataEntry](#containersai.datahub.prediction.v1alpha2.NodePrediction.PredictedRawDataEntry) | repeated | use containersai.datahub.metric.v1alpha2.MetricType as key |
| is_scheduled | [bool](#bool) |  |  |






<a name="containersai.datahub.prediction.v1alpha2.NodePrediction.PredictedRawDataEntry"></a>

### NodePrediction.PredictedRawDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [containersai.datahub.metric.v1alpha2.MetricData](#containersai.datahub.metric.v1alpha2.MetricData) |  |  |






<a name="containersai.datahub.prediction.v1alpha2.PodPrediction"></a>

### PodPrediction
Represents a list of predicted metrics data of a pod


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespaced_name | [containersai.datahub.resource.metadata.v1alpha2.NamespacedName](#containersai.datahub.resource.metadata.v1alpha2.NamespacedName) |  |  |
| container_predictions | [ContainerPrediction](#containersai.datahub.prediction.v1alpha2.ContainerPrediction) | repeated |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

