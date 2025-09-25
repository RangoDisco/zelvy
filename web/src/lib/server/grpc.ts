import {API_KEY} from "$env/static/private";
import {Metadata} from "@grpc/grpc-js";


export const createMetadataWithAuth = () => {
    const metadata = new Metadata();
    metadata.add("authorization", API_KEY);

    return metadata;
};
