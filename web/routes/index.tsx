import { Layout } from "../components/Layout.tsx";
import { Handlers } from "$fresh/server.ts";

type Character = {
    id: number;
    name: string;
    blood: string;
    species: string;
    patronus: string;
    born: string;
    quote: string;
    imgUrl: URL;
};

const api = "http://localhost:8080/api";

export const handler: Handlers<Character[] | null> = {
    async GET(_req, ctx) {
        const res = await fetch(api + "/characters");
        if (res.status === 404) {
            return ctx.render(null);
        }

        const characters = await res.json();
        return ctx.render(characters);
    }
};

export default function Home() {
    return (
        <Layout title="Welcome Home">
            <main>
                List of characters go here
            </main>
        </Layout>
    );
}
