import {inject, injectable} from "inversify";
import {UseCase} from "../interactor/UseCase";
import {CatalogItem} from "./CatalogItem";
import {Types} from "../di/types";
import {CatalogRepository} from "./CatalogRepository";
import {TaskEither} from "fp-ts/TaskEither";
import {pipe} from "fp-ts/function";
import {taskEither as TE} from "fp-ts";
import {CatalogError} from "./CatalogError";

@injectable()
export class GetCatalogItemsByTags extends UseCase<[string[], number, number], CatalogItem[]> {
    constructor(@inject(Types.catalog.repository) private readonly repository: CatalogRepository) {
        super();
    }

    execute = (arg: [string[], number, number]): TaskEither<Error, CatalogItem[]> => pipe(
        this.repository.getItemsByTags(...arg),
        TE.chain(TE.fromNullable(CatalogError.NotFound))
    );
}