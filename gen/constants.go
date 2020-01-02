package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  export(id: ID, q: String, filter: ExportFilterType): Export
  exports(offset: Int, limit: Int = 30, q: String, sort: [ExportSortType!], filter: ExportFilterType): ExportResultType
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
  errorDescription: String
  fileId: String
}

input ExportUpdateInput {
  type: String
  metadata: String
  state: ExportState
  errorDescription: String
  fileId: String
}

input ExportSortType {
  id: ObjectSortType
  type: ObjectSortType
  metadata: ObjectSortType
  state: ObjectSortType
  errorDescription: ObjectSortType
  fileId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
}

input ExportFilterType {
  AND: [ExportFilterType!]
  OR: [ExportFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  type: String
  type_ne: String
  type_gt: String
  type_lt: String
  type_gte: String
  type_lte: String
  type_in: [String!]
  type_like: String
  type_prefix: String
  type_suffix: String
  type_null: Boolean
  metadata: String
  metadata_ne: String
  metadata_gt: String
  metadata_lt: String
  metadata_gte: String
  metadata_lte: String
  metadata_in: [String!]
  metadata_like: String
  metadata_prefix: String
  metadata_suffix: String
  metadata_null: Boolean
  state: ExportState
  state_ne: ExportState
  state_gt: ExportState
  state_lt: ExportState
  state_gte: ExportState
  state_lte: ExportState
  state_in: [ExportState!]
  state_null: Boolean
  errorDescription: String
  errorDescription_ne: String
  errorDescription_gt: String
  errorDescription_lt: String
  errorDescription_gte: String
  errorDescription_lte: String
  errorDescription_in: [String!]
  errorDescription_like: String
  errorDescription_prefix: String
  errorDescription_suffix: String
  errorDescription_null: Boolean
  fileId: String
  fileId_ne: String
  fileId_gt: String
  fileId_lt: String
  fileId_gte: String
  fileId_lte: String
  fileId_in: [String!]
  fileId_like: String
  fileId_prefix: String
  fileId_suffix: String
  fileId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
}

type ExportResultType {
  items: [Export!]!
  count: Int!
}`
)
