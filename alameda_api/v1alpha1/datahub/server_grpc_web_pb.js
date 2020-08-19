/**
 * @fileoverview gRPC-Web generated client stub for containersai.alameda.v1alpha1.datahub
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var alameda_api_v1alpha1_datahub_applications_services_pb = require('../../../alameda_api/v1alpha1/datahub/applications/services_pb.js')

var alameda_api_v1alpha1_datahub_data_services_pb = require('../../../alameda_api/v1alpha1/datahub/data/services_pb.js')

var alameda_api_v1alpha1_datahub_events_services_pb = require('../../../alameda_api/v1alpha1/datahub/events/services_pb.js')

var alameda_api_v1alpha1_datahub_gpu_services_pb = require('../../../alameda_api/v1alpha1/datahub/gpu/services_pb.js')

var alameda_api_v1alpha1_datahub_licenses_services_pb = require('../../../alameda_api/v1alpha1/datahub/licenses/services_pb.js')

var alameda_api_v1alpha1_datahub_metrics_services_pb = require('../../../alameda_api/v1alpha1/datahub/metrics/services_pb.js')

var alameda_api_v1alpha1_datahub_plannings_services_pb = require('../../../alameda_api/v1alpha1/datahub/plannings/services_pb.js')

var alameda_api_v1alpha1_datahub_predictions_services_pb = require('../../../alameda_api/v1alpha1/datahub/predictions/services_pb.js')

var alameda_api_v1alpha1_datahub_rawdata_services_pb = require('../../../alameda_api/v1alpha1/datahub/rawdata/services_pb.js')

var alameda_api_v1alpha1_datahub_recommendations_services_pb = require('../../../alameda_api/v1alpha1/datahub/recommendations/services_pb.js')

var alameda_api_v1alpha1_datahub_resources_services_pb = require('../../../alameda_api/v1alpha1/datahub/resources/services_pb.js')

var alameda_api_v1alpha1_datahub_schemas_services_pb = require('../../../alameda_api/v1alpha1/datahub/schemas/services_pb.js')

var alameda_api_v1alpha1_datahub_scores_services_pb = require('../../../alameda_api/v1alpha1/datahub/scores/services_pb.js')

var alameda_api_v1alpha1_datahub_weavescope_services_pb = require('../../../alameda_api/v1alpha1/datahub/weavescope/services_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_rpc_status_pb = require('../../../google/rpc/status_pb.js')
const proto = {};
proto.containersai = {};
proto.containersai.alameda = {};
proto.containersai.alameda.v1alpha1 = {};
proto.containersai.alameda.v1alpha1.datahub = require('./server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApps = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApps',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_applications_services_pb.CreateApplicationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApps = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApps,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApps);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse>}
 */
const methodDescriptor_DatahubService_ListApps = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApps',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_applications_services_pb.ListApplicationsRequest,
  alameda_api_v1alpha1_datahub_applications_services_pb.ListApplicationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_applications_services_pb.ListApplicationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse>}
 */
const methodInfo_DatahubService_ListApps = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_applications_services_pb.ListApplicationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_applications_services_pb.ListApplicationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApps,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApps);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteApps = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApps',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_applications_services_pb.DeleteApplicationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteApps = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteApps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteApps,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteApps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApps',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteApps);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.data.ReadDataResponse>}
 */
const methodDescriptor_DatahubService_ReadData = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadData',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_data_services_pb.ReadDataRequest,
  alameda_api_v1alpha1_datahub_data_services_pb.ReadDataResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_data_services_pb.ReadDataResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.data.ReadDataResponse>}
 */
const methodInfo_DatahubService_ReadData = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_data_services_pb.ReadDataResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_data_services_pb.ReadDataResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.data.ReadDataResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.readData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ReadData,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.data.ReadDataResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.readData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ReadData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_WriteData = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteData',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_data_services_pb.WriteDataRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_WriteData = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.writeData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_WriteData,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.WriteDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.writeData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_WriteData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteData = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteData',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_data_services_pb.DeleteDataRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteData = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteData,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteData',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateEvents = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateEvents',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_events_services_pb.CreateEventsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateEvents = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateEvents',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateEvents,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.CreateEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateEvents',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.events.ListEventsResponse>}
 */
const methodDescriptor_DatahubService_ListEvents = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListEvents',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_events_services_pb.ListEventsRequest,
  alameda_api_v1alpha1_datahub_events_services_pb.ListEventsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_events_services_pb.ListEventsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.events.ListEventsResponse>}
 */
const methodInfo_DatahubService_ListEvents = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_events_services_pb.ListEventsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_events_services_pb.ListEventsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.events.ListEventsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListEvents',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListEvents,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.events.ListEventsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListEvents',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateGpuPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateGpuPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_gpu_services_pb.CreateGpuPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateGpuPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createGpuPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateGpuPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateGpuPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createGpuPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateGpuPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateGpuPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse>}
 */
const methodDescriptor_DatahubService_ListGpus = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpus',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpusRequest,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpusResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpusResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse>}
 */
const methodInfo_DatahubService_ListGpus = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpusResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpusResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listGpus =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpus',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpus,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listGpus =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpus',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpus);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListGpuMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuMetricsRequest,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse>}
 */
const methodInfo_DatahubService_ListGpuMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listGpuMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpuMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listGpuMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpuMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListGpuPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuPredictionsRequest,
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse>}
 */
const methodInfo_DatahubService_ListGpuPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_gpu_services_pb.ListGpuPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listGpuPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpuPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listGpuPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListGpuPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListGpuPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse>}
 */
const methodDescriptor_DatahubService_GetLicense = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/GetLicense',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  alameda_api_v1alpha1_datahub_licenses_services_pb.GetLicenseResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_licenses_services_pb.GetLicenseResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse>}
 */
const methodInfo_DatahubService_GetLicense = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_licenses_services_pb.GetLicenseResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_licenses_services_pb.GetLicenseResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.getLicense =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetLicense',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetLicense,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.licenses.GetLicenseResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.getLicense =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetLicense',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetLicense);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreatePodMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPodMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreatePodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPodMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateControllerMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateControllerMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateControllerMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createControllerMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateControllerMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createControllerMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApplicationMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateApplicationMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApplicationMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApplicationMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateApplicationMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApplicationMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNamespaceMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateNamespaceMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNamespaceMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNamespaceMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaceMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNamespaceMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNamespaceMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaceMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodeMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateNodeMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodeMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNodeMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodeMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNodeMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodeMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateClusterMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.CreateClusterMetricsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateClusterMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createClusterMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.CreateClusterMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createClusterMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse>}
 */
const methodInfo_DatahubService_ListMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListPodMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListPodMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListPodMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListPodMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse>}
 */
const methodInfo_DatahubService_ListPodMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListPodMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListPodMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPodMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListPodMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPodMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListControllerMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListControllerMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListControllerMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListControllerMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse>}
 */
const methodInfo_DatahubService_ListControllerMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListControllerMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListControllerMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listControllerMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListControllerMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listControllerMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListApplicationMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListApplicationMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListApplicationMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListApplicationMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse>}
 */
const methodInfo_DatahubService_ListApplicationMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListApplicationMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListApplicationMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApplicationMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListApplicationMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApplicationMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListNamespaceMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNamespaceMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNamespaceMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNamespaceMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse>}
 */
const methodInfo_DatahubService_ListNamespaceMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNamespaceMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNamespaceMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNamespaceMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaceMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNamespaceMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNamespaceMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaceMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListNodeMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNodeMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNodeMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNodeMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse>}
 */
const methodInfo_DatahubService_ListNodeMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNodeMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListNodeMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNodeMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListNodeMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNodeMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListClusterMetrics = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterMetrics',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListClusterMetricsRequest,
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListClusterMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListClusterMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse>}
 */
const methodInfo_DatahubService_ListClusterMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListClusterMetricsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_metrics_services_pb.ListClusterMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listClusterMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterMetrics,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.metrics.ListClusterMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listClusterMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_Ping = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/Ping',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_Ping = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.ping =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/Ping',
      request,
      metadata || {},
      methodDescriptor_DatahubService_Ping,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.ping =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/Ping',
      request,
      metadata || {},
      methodDescriptor_DatahubService_Ping);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreatePlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreatePodPlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPodPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreatePodPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPodPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateControllerPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreateControllerPlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateControllerPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createControllerPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateControllerPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createControllerPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApplicationPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreateApplicationPlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApplicationPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApplicationPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateApplicationPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApplicationPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNamespacePlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreateNamespacePlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNamespacePlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNamespacePlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespacePlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNamespacePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNamespacePlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespacePlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodePlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreateNodePlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodePlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNodePlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateNodePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNodePlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateClusterPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.CreateClusterPlanningsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateClusterPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createClusterPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.CreateClusterPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createClusterPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse>}
 */
const methodInfo_DatahubService_ListPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListPodPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPodPlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPodPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPodPlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse>}
 */
const methodInfo_DatahubService_ListPodPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPodPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListPodPlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPodPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListPodPlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPodPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListControllerPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListControllerPlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListControllerPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListControllerPlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse>}
 */
const methodInfo_DatahubService_ListControllerPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListControllerPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListControllerPlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listControllerPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListControllerPlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listControllerPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListApplicationPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListApplicationPlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListApplicationPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListApplicationPlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse>}
 */
const methodInfo_DatahubService_ListApplicationPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListApplicationPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListApplicationPlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApplicationPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListApplicationPlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApplicationPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListNamespacePlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNamespacePlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNamespacePlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNamespacePlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse>}
 */
const methodInfo_DatahubService_ListNamespacePlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNamespacePlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNamespacePlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNamespacePlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespacePlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNamespacePlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNamespacePlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespacePlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListNodePlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNodePlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNodePlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNodePlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse>}
 */
const methodInfo_DatahubService_ListNodePlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNodePlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListNodePlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNodePlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListNodePlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNodePlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse>}
 */
const methodDescriptor_DatahubService_ListClusterPlannings = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPlannings',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListClusterPlanningsRequest,
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListClusterPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListClusterPlanningsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse>}
 */
const methodInfo_DatahubService_ListClusterPlannings = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListClusterPlanningsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_plannings_services_pb.ListClusterPlanningsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listClusterPlannings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterPlannings,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.plannings.ListClusterPlanningsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listClusterPlannings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPlannings',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterPlannings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreatePredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreatePodPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPodPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreatePodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPodPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateControllerPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreateControllerPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateControllerPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createControllerPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateControllerPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createControllerPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApplicationPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreateApplicationPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApplicationPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApplicationPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateApplicationPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApplicationPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNamespacePredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreateNamespacePredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNamespacePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNamespacePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespacePredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNamespacePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNamespacePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespacePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespacePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodePredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreateNodePredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNodePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNodePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateClusterPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.CreateClusterPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateClusterPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createClusterPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.CreateClusterPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createClusterPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse>}
 */
const methodInfo_DatahubService_ListPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListPodPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPodPredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPodPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPodPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse>}
 */
const methodInfo_DatahubService_ListPodPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPodPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListPodPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPodPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListPodPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPodPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListControllerPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListControllerPredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListControllerPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListControllerPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse>}
 */
const methodInfo_DatahubService_ListControllerPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListControllerPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListControllerPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listControllerPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListControllerPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listControllerPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListApplicationPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListApplicationPredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListApplicationPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListApplicationPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse>}
 */
const methodInfo_DatahubService_ListApplicationPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListApplicationPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListApplicationPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApplicationPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListApplicationPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApplicationPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListNamespacePredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNamespacePredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNamespacePredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNamespacePredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse>}
 */
const methodInfo_DatahubService_ListNamespacePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNamespacePredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNamespacePredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNamespacePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespacePredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNamespacePredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNamespacePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespacePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespacePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListNodePredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNodePredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNodePredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNodePredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse>}
 */
const methodInfo_DatahubService_ListNodePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNodePredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListNodePredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNodePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListNodePredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNodePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListClusterPredictions = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPredictions',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListClusterPredictionsRequest,
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListClusterPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListClusterPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse>}
 */
const methodInfo_DatahubService_ListClusterPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListClusterPredictionsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_predictions_services_pb.ListClusterPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listClusterPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterPredictions,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.predictions.ListClusterPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listClusterPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse>}
 */
const methodDescriptor_DatahubService_ReadRawdata = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadRawdata',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_rawdata_services_pb.ReadRawdataRequest,
  alameda_api_v1alpha1_datahub_rawdata_services_pb.ReadRawdataResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_rawdata_services_pb.ReadRawdataResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse>}
 */
const methodInfo_DatahubService_ReadRawdata = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_rawdata_services_pb.ReadRawdataResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_rawdata_services_pb.ReadRawdataResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.readRawdata =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadRawdata',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ReadRawdata,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.readRawdata =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ReadRawdata',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ReadRawdata);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_WriteRawdata = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteRawdata',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_rawdata_services_pb.WriteRawdataRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_WriteRawdata = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.writeRawdata =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteRawdata',
      request,
      metadata || {},
      methodDescriptor_DatahubService_WriteRawdata,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.writeRawdata =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/WriteRawdata',
      request,
      metadata || {},
      methodDescriptor_DatahubService_WriteRawdata);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreatePodRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPodRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPodRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateControllerRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateControllerRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateControllerRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createControllerRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createControllerRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllerRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllerRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApplicationRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateApplicationRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApplicationRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApplicationRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateApplicationRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApplicationRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplicationRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplicationRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNamespaceRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateNamespaceRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNamespaceRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNamespaceRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaceRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNamespaceRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNamespaceRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaceRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaceRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodeRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateNodeRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodeRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNodeRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodeRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateNodeRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNodeRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodeRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodeRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateClusterRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.CreateClusterRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateClusterRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createClusterRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.CreateClusterRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createClusterRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusterRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusterRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListPodRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListPodRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPodRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPodRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListAvailablePodRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListAvailablePodRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListAvailablePodRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listAvailablePodRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListAvailablePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListAvailablePodRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listAvailablePodRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListAvailablePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListAvailablePodRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListControllerRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListControllerRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListControllerRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListControllerRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListControllerRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListControllerRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListControllerRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listControllerRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listControllerRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllerRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllerRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListApplicationRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListApplicationRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListApplicationRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListApplicationRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApplicationRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListApplicationRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApplicationRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplicationRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplicationRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListNamespaceRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNamespaceRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListNamespaceRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNamespaceRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNamespaceRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaceRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNamespaceRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNamespaceRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaceRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaceRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListNodeRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNodeRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNodeRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNodeRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListNodeRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNodeRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListNodeRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNodeRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListNodeRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNodeRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodeRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListClusterRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterRecommendations',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListClusterRecommendationsRequest,
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListClusterRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListClusterRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListClusterRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListClusterRecommendationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_recommendations_services_pb.ListClusterRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listClusterRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.recommendations.ListClusterRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listClusterRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusterRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusterRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePods = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePods',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreatePodsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePods = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createPods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePods,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createPods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateControllers = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllers',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreateControllersRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateControllers = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createControllers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllers,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createControllers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateControllers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateApplications = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplications',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreateApplicationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateApplications = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createApplications =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplications,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createApplications =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateApplications);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNamespaces = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaces',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreateNamespacesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNamespaces = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNamespaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaces,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNamespaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNamespaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodes = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodes',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreateNodesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodes = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodes,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateClusters = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusters',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.CreateClustersRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateClusters = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createClusters =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusters,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.CreateClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createClusters =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateClusters);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse>}
 */
const methodDescriptor_DatahubService_ListPods = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPods',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListPodsRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListPodsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListPodsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse>}
 */
const methodInfo_DatahubService_ListPods = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListPodsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListPodsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listPods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPods,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListPodsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listPods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListPods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse>}
 */
const methodDescriptor_DatahubService_ListControllers = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllers',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListControllersRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListControllersResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListControllersResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse>}
 */
const methodInfo_DatahubService_ListControllers = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListControllersResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListControllersResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listControllers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllers,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListControllersResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listControllers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListControllers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse>}
 */
const methodDescriptor_DatahubService_ListApplications = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplications',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListApplicationsRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListApplicationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListApplicationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse>}
 */
const methodInfo_DatahubService_ListApplications = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListApplicationsResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListApplicationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listApplications =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplications,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListApplicationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listApplications =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListApplications);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse>}
 */
const methodDescriptor_DatahubService_ListNamespaces = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaces',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNamespacesRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNamespacesResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNamespacesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse>}
 */
const methodInfo_DatahubService_ListNamespaces = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNamespacesResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNamespacesResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNamespaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaces,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListNamespacesResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNamespaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNamespaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse>}
 */
const methodDescriptor_DatahubService_ListNodes = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodes',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNodesRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNodesResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNodesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse>}
 */
const methodInfo_DatahubService_ListNodes = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNodesResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListNodesResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodes,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListNodesResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse>}
 */
const methodDescriptor_DatahubService_ListClusters = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusters',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListClustersRequest,
  alameda_api_v1alpha1_datahub_resources_services_pb.ListClustersResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListClustersResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse>}
 */
const methodInfo_DatahubService_ListClusters = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_resources_services_pb.ListClustersResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_resources_services_pb.ListClustersResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listClusters =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusters,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.resources.ListClustersResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listClusters =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListClusters);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeletePods = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeletePods',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeletePodsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeletePods = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deletePods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeletePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeletePods,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeletePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deletePods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeletePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeletePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteControllers = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteControllers',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeleteControllersRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteControllers = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteControllers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteControllers,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteControllersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteControllers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteControllers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteControllers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteApplications = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApplications',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeleteApplicationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteApplications = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteApplications =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteApplications,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteApplicationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteApplications =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteApplications',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteApplications);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteNamespaces = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNamespaces',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeleteNamespacesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteNamespaces = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteNamespaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNamespaces,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteNamespaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNamespaces',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNamespaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteNodes = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNodes',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeleteNodesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteNodes = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNodes,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteClusters = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteClusters',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_resources_services_pb.DeleteClustersRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteClusters = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteClusters =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteClusters,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.resources.DeleteClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteClusters =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteClusters',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteClusters);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateSchemas = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSchemas',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_schemas_services_pb.CreateSchemasRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateSchemas = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createSchemas =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSchemas,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createSchemas =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSchemas);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse>}
 */
const methodDescriptor_DatahubService_ListSchemas = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSchemas',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_schemas_services_pb.ListSchemasRequest,
  alameda_api_v1alpha1_datahub_schemas_services_pb.ListSchemasResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_schemas_services_pb.ListSchemasResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse>}
 */
const methodInfo_DatahubService_ListSchemas = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_schemas_services_pb.ListSchemasResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_schemas_services_pb.ListSchemasResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listSchemas =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSchemas,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listSchemas =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSchemas);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteSchemas = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteSchemas',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_schemas_services_pb.DeleteSchemasRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteSchemas = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.deleteSchemas =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteSchemas,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.deleteSchemas =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/DeleteSchemas',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteSchemas);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateSimulatedSchedulingScores = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSimulatedSchedulingScores',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_scores_services_pb.CreateSimulatedSchedulingScoresRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateSimulatedSchedulingScores = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.createSimulatedSchedulingScores =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSimulatedSchedulingScores,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.createSimulatedSchedulingScores =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/CreateSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSimulatedSchedulingScores);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse>}
 */
const methodDescriptor_DatahubService_ListSimulatedSchedulingScores = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSimulatedSchedulingScores',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_scores_services_pb.ListSimulatedSchedulingScoresRequest,
  alameda_api_v1alpha1_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse>}
 */
const methodInfo_DatahubService_ListSimulatedSchedulingScores = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_scores_services_pb.ListSimulatedSchedulingScoresResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listSimulatedSchedulingScores =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSimulatedSchedulingScores,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listSimulatedSchedulingScores =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSimulatedSchedulingScores);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_ListWeaveScopeHosts = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeHosts',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_ListWeaveScopeHosts = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listWeaveScopeHosts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeHosts',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeHosts,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listWeaveScopeHosts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeHosts',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeHosts);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_GetWeaveScopeHostDetails = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeHostDetails',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeHostsRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_GetWeaveScopeHostDetails = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.getWeaveScopeHostDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeHostDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopeHostDetails,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.getWeaveScopeHostDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeHostDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopeHostDetails);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_ListWeaveScopePods = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopePods',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopePodsRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_ListWeaveScopePods = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listWeaveScopePods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopePods,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listWeaveScopePods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_GetWeaveScopePodDetails = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopePodDetails',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopePodsRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_GetWeaveScopePodDetails = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.getWeaveScopePodDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopePodDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopePodDetails,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.getWeaveScopePodDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopePodDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopePodDetails);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_ListWeaveScopeContainers = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainers',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_ListWeaveScopeContainers = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listWeaveScopeContainers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainers,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listWeaveScopeContainers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainers',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_ListWeaveScopeContainersByHostname = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByHostname',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_ListWeaveScopeContainersByHostname = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listWeaveScopeContainersByHostname =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByHostname',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainersByHostname,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listWeaveScopeContainersByHostname =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByHostname',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainersByHostname);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_ListWeaveScopeContainersByImage = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByImage',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_ListWeaveScopeContainersByImage = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.listWeaveScopeContainersByImage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByImage',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainersByImage,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.listWeaveScopeContainersByImage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/ListWeaveScopeContainersByImage',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListWeaveScopeContainersByImage);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodDescriptor_DatahubService_GetWeaveScopeContainerDetails = new grpc.web.MethodDescriptor(
  '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeContainerDetails',
  grpc.web.MethodType.UNARY,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.ListWeaveScopeContainersRequest,
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest,
 *   !proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 */
const methodInfo_DatahubService_GetWeaveScopeContainerDetails = new grpc.web.AbstractClientBase.MethodInfo(
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse,
  /**
   * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  alameda_api_v1alpha1_datahub_weavescope_services_pb.WeaveScopeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServiceClient.prototype.getWeaveScopeContainerDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeContainerDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopeContainerDetails,
      callback);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.alameda.v1alpha1.datahub.DatahubServicePromiseClient.prototype.getWeaveScopeContainerDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.alameda.v1alpha1.datahub.DatahubService/GetWeaveScopeContainerDetails',
      request,
      metadata || {},
      methodDescriptor_DatahubService_GetWeaveScopeContainerDetails);
};


module.exports = proto.containersai.alameda.v1alpha1.datahub;

