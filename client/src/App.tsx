import { useEffect, useState } from 'react';

type Person = {
  Name: string
}

export default function App() {
  const [goRes, setGores] = useState<Person>({Name: ""});

  async function fetchGo() {
    const res = await fetch('http://localhost:8080/person');
    setGores(await res.json());
  }

  useEffect(() => {
    fetchGo();
  }, [])

  return(
    <h1>{goRes.Name}</h1>
  );
}