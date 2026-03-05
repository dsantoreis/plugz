import React, { useEffect, useState } from 'react'
import { createRoot } from 'react-dom/client'
import './styles.css'

const API = import.meta.env.VITE_API_URL || 'http://localhost:8080'
const TOKEN = import.meta.env.VITE_API_TOKEN || 'dev-token'

function App() {
  const [skills, setSkills] = useState([])
  const [result, setResult] = useState('')

  async function loadCatalog() {
    const res = await fetch(`${API}/api/v1/catalog`, {
      headers: { Authorization: `Bearer ${TOKEN}` }
    })
    setSkills(await res.json())
  }

  async function install(name) {
    const res = await fetch(`${API}/api/v1/install`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${TOKEN}`
      },
      body: JSON.stringify({ name })
    })
    setResult(JSON.stringify(await res.json(), null, 2))
  }

  async function testSkill(name) {
    const res = await fetch(`${API}/api/v1/test/${name}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${TOKEN}`
      },
      body: JSON.stringify({ input: 'hello marketplace' })
    })
    setResult(JSON.stringify(await res.json(), null, 2))
  }

  useEffect(() => { loadCatalog() }, [])

  return (
    <main>
      <h1>AI Skill Marketplace</h1>
      <p>Catalog + install + test flow</p>
      <div className='grid'>
        {skills.map(s => (
          <article key={s.name}>
            <h3>{s.name}</h3>
            <p>{s.description || 'No description'}</p>
            <div className='row'>
              <button onClick={() => install(s.name)}>Install</button>
              <button onClick={() => testSkill(s.name)}>Test</button>
            </div>
          </article>
        ))}
      </div>
      <pre>{result}</pre>
    </main>
  )
}

createRoot(document.getElementById('root')).render(<App />)
