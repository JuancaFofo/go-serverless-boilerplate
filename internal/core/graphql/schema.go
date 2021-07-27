package coregraphql

const AppSchema = `
	type TodoItem {
		id: Int!
		description: String!
		completed: Boolean!
	}
	input TodoItemInput {
		description: String!
		completed: Boolean!
	}

	type Query {
		hello: String!
	}
	type Mutation {
		addTodo(todo: TodoItemInput!): TodoItem
	}
`
