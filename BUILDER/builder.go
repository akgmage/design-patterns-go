// Some objects are simple and can be created in a single constructor call
// Other objeccts require a lot of ceremony to create
// Having a factory function with 15 arguments is not productive
// Instead, opt for piecewise (piece-by-piece) construction
// BUilder provides an API for constructing an object step by step