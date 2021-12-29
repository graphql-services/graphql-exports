package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyHTTPRequest         key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  export(id: ID, q: String, filter: ExportFilterType): Export
  exports(offset: Int, limit: Int = 30, q: String, sort: [ExportSortType!], filter: ExportFilterType): ExportResultType!
}

type Mutation {
  createExport(input: ExportCreateInput!): Export!
  updateExport(id: ID!, input: ExportUpdateInput!): Export!
  deleteExport(id: ID!): Export!
  deleteAllExports: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

enum ExportState {
  PROCESSING
  COMPLETED
  ERROR
}

type Export {
  id: ID!
  type: String
  metadata: String
  state: ExportState
  progress: Float
  errorDescription: String
  fileId: String
  file: File
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

extend type File @key(fields: "id") {
  id: ID! @external
}

input ExportCreateInput {
  id: ID
  type: String
  metadata: String
  state: ExportState
  progress: Float
  errorDescription: String
  fileId: String
}

input ExportUpdateInput {
  type: String
  metadata: String
  state: ExportState
  progress: Float
  errorDescription: String
  fileId: String
}

input ExportSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  type: ObjectSortType
  typeMin: ObjectSortType
  typeMax: ObjectSortType
  metadata: ObjectSortType
  metadataMin: ObjectSortType
  metadataMax: ObjectSortType
  state: ObjectSortType
  stateMin: ObjectSortType
  stateMax: ObjectSortType
  progress: ObjectSortType
  progressMin: ObjectSortType
  progressMax: ObjectSortType
  progressAvg: ObjectSortType
  progressSum: ObjectSortType
  errorDescription: ObjectSortType
  errorDescriptionMin: ObjectSortType
  errorDescriptionMax: ObjectSortType
  fileId: ObjectSortType
  fileIdMin: ObjectSortType
  fileIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
}

input ExportFilterType {
  AND: [ExportFilterType!]
  OR: [ExportFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_not_in: [ID!]
  idMin_not_in: [ID!]
  idMax_not_in: [ID!]
  id_null: Boolean
  type: String
  typeMin: String
  typeMax: String
  type_ne: String
  typeMin_ne: String
  typeMax_ne: String
  type_gt: String
  typeMin_gt: String
  typeMax_gt: String
  type_lt: String
  typeMin_lt: String
  typeMax_lt: String
  type_gte: String
  typeMin_gte: String
  typeMax_gte: String
  type_lte: String
  typeMin_lte: String
  typeMax_lte: String
  type_in: [String!]
  typeMin_in: [String!]
  typeMax_in: [String!]
  type_not_in: [String!]
  typeMin_not_in: [String!]
  typeMax_not_in: [String!]
  type_like: String
  typeMin_like: String
  typeMax_like: String
  type_prefix: String
  typeMin_prefix: String
  typeMax_prefix: String
  type_suffix: String
  typeMin_suffix: String
  typeMax_suffix: String
  type_null: Boolean
  metadata: String
  metadataMin: String
  metadataMax: String
  metadata_ne: String
  metadataMin_ne: String
  metadataMax_ne: String
  metadata_gt: String
  metadataMin_gt: String
  metadataMax_gt: String
  metadata_lt: String
  metadataMin_lt: String
  metadataMax_lt: String
  metadata_gte: String
  metadataMin_gte: String
  metadataMax_gte: String
  metadata_lte: String
  metadataMin_lte: String
  metadataMax_lte: String
  metadata_in: [String!]
  metadataMin_in: [String!]
  metadataMax_in: [String!]
  metadata_not_in: [String!]
  metadataMin_not_in: [String!]
  metadataMax_not_in: [String!]
  metadata_like: String
  metadataMin_like: String
  metadataMax_like: String
  metadata_prefix: String
  metadataMin_prefix: String
  metadataMax_prefix: String
  metadata_suffix: String
  metadataMin_suffix: String
  metadataMax_suffix: String
  metadata_null: Boolean
  state: ExportState
  stateMin: ExportState
  stateMax: ExportState
  state_ne: ExportState
  stateMin_ne: ExportState
  stateMax_ne: ExportState
  state_gt: ExportState
  stateMin_gt: ExportState
  stateMax_gt: ExportState
  state_lt: ExportState
  stateMin_lt: ExportState
  stateMax_lt: ExportState
  state_gte: ExportState
  stateMin_gte: ExportState
  stateMax_gte: ExportState
  state_lte: ExportState
  stateMin_lte: ExportState
  stateMax_lte: ExportState
  state_in: [ExportState!]
  stateMin_in: [ExportState!]
  stateMax_in: [ExportState!]
  state_not_in: [ExportState!]
  stateMin_not_in: [ExportState!]
  stateMax_not_in: [ExportState!]
  state_null: Boolean
  progress: Float
  progressMin: Float
  progressMax: Float
  progressAvg: Float
  progressSum: Float
  progress_ne: Float
  progressMin_ne: Float
  progressMax_ne: Float
  progressAvg_ne: Float
  progressSum_ne: Float
  progress_gt: Float
  progressMin_gt: Float
  progressMax_gt: Float
  progressAvg_gt: Float
  progressSum_gt: Float
  progress_lt: Float
  progressMin_lt: Float
  progressMax_lt: Float
  progressAvg_lt: Float
  progressSum_lt: Float
  progress_gte: Float
  progressMin_gte: Float
  progressMax_gte: Float
  progressAvg_gte: Float
  progressSum_gte: Float
  progress_lte: Float
  progressMin_lte: Float
  progressMax_lte: Float
  progressAvg_lte: Float
  progressSum_lte: Float
  progress_in: [Float!]
  progressMin_in: [Float!]
  progressMax_in: [Float!]
  progressAvg_in: [Float!]
  progressSum_in: [Float!]
  progress_not_in: [Float!]
  progressMin_not_in: [Float!]
  progressMax_not_in: [Float!]
  progressAvg_not_in: [Float!]
  progressSum_not_in: [Float!]
  progress_null: Boolean
  errorDescription: String
  errorDescriptionMin: String
  errorDescriptionMax: String
  errorDescription_ne: String
  errorDescriptionMin_ne: String
  errorDescriptionMax_ne: String
  errorDescription_gt: String
  errorDescriptionMin_gt: String
  errorDescriptionMax_gt: String
  errorDescription_lt: String
  errorDescriptionMin_lt: String
  errorDescriptionMax_lt: String
  errorDescription_gte: String
  errorDescriptionMin_gte: String
  errorDescriptionMax_gte: String
  errorDescription_lte: String
  errorDescriptionMin_lte: String
  errorDescriptionMax_lte: String
  errorDescription_in: [String!]
  errorDescriptionMin_in: [String!]
  errorDescriptionMax_in: [String!]
  errorDescription_not_in: [String!]
  errorDescriptionMin_not_in: [String!]
  errorDescriptionMax_not_in: [String!]
  errorDescription_like: String
  errorDescriptionMin_like: String
  errorDescriptionMax_like: String
  errorDescription_prefix: String
  errorDescriptionMin_prefix: String
  errorDescriptionMax_prefix: String
  errorDescription_suffix: String
  errorDescriptionMin_suffix: String
  errorDescriptionMax_suffix: String
  errorDescription_null: Boolean
  fileId: String
  fileIdMin: String
  fileIdMax: String
  fileId_ne: String
  fileIdMin_ne: String
  fileIdMax_ne: String
  fileId_gt: String
  fileIdMin_gt: String
  fileIdMax_gt: String
  fileId_lt: String
  fileIdMin_lt: String
  fileIdMax_lt: String
  fileId_gte: String
  fileIdMin_gte: String
  fileIdMax_gte: String
  fileId_lte: String
  fileIdMin_lte: String
  fileIdMax_lte: String
  fileId_in: [String!]
  fileIdMin_in: [String!]
  fileIdMax_in: [String!]
  fileId_not_in: [String!]
  fileIdMin_not_in: [String!]
  fileIdMax_not_in: [String!]
  fileId_like: String
  fileIdMin_like: String
  fileIdMax_like: String
  fileId_prefix: String
  fileIdMin_prefix: String
  fileIdMax_prefix: String
  fileId_suffix: String
  fileIdMin_suffix: String
  fileIdMax_suffix: String
  fileId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_not_in: [Time!]
  updatedAtMin_not_in: [Time!]
  updatedAtMax_not_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_not_in: [Time!]
  createdAtMin_not_in: [Time!]
  createdAtMax_not_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_not_in: [ID!]
  updatedByMin_not_in: [ID!]
  updatedByMax_not_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_not_in: [ID!]
  createdByMin_not_in: [ID!]
  createdByMax_not_in: [ID!]
  createdBy_null: Boolean
}

type ExportResultType {
  items: [Export!]!
  count: Int!
  aggregations: ExportResultAggregations!
}

type ExportResultAggregations {
  typeMin: String
  typeMax: String
  metadataMin: String
  metadataMax: String
  progressMin: Float
  progressMax: Float
  progressAvg: Float
  progressSum: Float
  errorDescriptionMin: String
  errorDescriptionMax: String
  fileIdMin: String
  fileIdMax: String
}`
)
