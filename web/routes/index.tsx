import { Layout } from "../components/Layout.tsx";
import { Handlers, PageProps } from "$fresh/server.ts";

const api = "http://localhost:8080/api";

export const handler: Handlers<Character[] | null> = {
    async GET(_req, ctx) {
        const res = await fetch(api + "/characters");
        if (!res.ok) {
            return ctx.render(null);
        }

        const characters: Character[] = await res.json();

        return ctx.render(characters);
    }
};

export default function Home({ data }: PageProps<Character[] | null>) {
    if (!data) return (
        <h1>Nothing to see</h1>
    );


    return (
        <Layout title="Welcome Home">
            <main>
                List of characters go here
                <List characters={data} />
            </main>
        </Layout>
    );
}

function List({ characters }: { characters: Character[]; }) {
    return (
        <section>
            {characters.map((character) => (
                <Card key={character.id} character={character} />
            ))}
        </section>
    );
}

type Character = {
    id: number;
    name: string;
    blood: string;
    species: string;
    patronus: string;
    born: string;
    quote: string;
    imgUrl: string;
};

type CardProp =
    & { character: Character; };

function Card({ character }: CardProp) {
    return (
        <div class="card">
            <h3>{character.name}</h3>
            <img src={character.imgUrl} alt="" />
            <dl>
                <dt>Blood type</dt>
                <dd>{character.blood}</dd>
                <dt>Species</dt>
                <dd>{character.species}</dd>
            </dl>
        </div>
    );
}
