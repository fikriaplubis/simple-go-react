import { useState, useEffect } from 'react'

export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <div>
      {error && <div>Failed to load {error.toString()}</div>}
      {!data ? <div> Loading... </div> : ((data?.data ?? []).length === 0 && <p>data kosong</p>)}
      <Input onSuccess={getData} />
      {data?.data && data?.data?.map((item, index) => (
        <div key={index}>
          <input type="checkbox" defaultChecked={item.done} disabled={true} />
          <span >ID: {item.ID} task: {item.task}</span>
          <Checklist item={item} onSuccess={getData} />
          <Delete id={item.ID} onSuccess={getData} />
        </div>
      ))}
    </div>
  )
}

function Checklist({ item, onSuccess }) {
  const [setError] = useState(null);

  const onClick = async (e) => {
    try {
      const body = {
        task: item.task,
        done: item.done == false ? true : false
      }

      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${item.ID}`, {
        method: 'PUT',
        body: JSON.stringify(body)
      });
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <button disabled={item.done == true ? true : false} onClick={onClick}>Done</button>
  )
}

function Delete({ id, onSuccess }) {
  const [setError] = useState(null);

  const onClick = async (e) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${id}`, {
        method: 'DELETE',
      });
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <button onClick={onClick}>Delete</button>
  )
}

function Input({ onSuccess }) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      task: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();
      setData(data.message);
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <div>
      {error && <p>error: {error.toString()}</p>}
      {data && <p>success: {data}</p>}
      <form onSubmit={handleSubmit}>
        <input name="data" type="text"/>
        <button >Submit</button>
      </form>
    </div>
  )
}