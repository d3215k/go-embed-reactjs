import { createFileRoute } from '@tanstack/react-router'

async function fetchHello() {
  return fetch("http://localhost:3000/hello")
    .then((r) => r.json());
}

export const Route = createFileRoute('/profile')({
  component: () => <ShowHello/>,
  loader: () => fetchHello(),
})

interface Hello {
  message: string
}

function ShowHello() {
  const hello: Hello = Route.useLoaderData();

  return <div>{hello.message}</div>
}