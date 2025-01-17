export namespace Types {
    export const app = {
        config: Symbol.for("config"),
        database: Symbol.for("database"),
        server: Symbol.for("server")
    };
    export const catalog = {
        collection: Symbol.for("catalogCollection"),
        repository: Symbol.for("catalogRepository"),
        service: Symbol.for("catalogService"),
        addCatalogItem: Symbol.for("addCatalogItem"),
        getCatalogItemById: Symbol.for("getCatalogItemById"),
        getCatalogItemsByTags: Symbol.for("getCatalogItemsByTags"),
        updateCatalogItem: Symbol.for("updateCatalogItem"),
        removeCatalogItem: Symbol.for("removeCatalogItem")
    };
}