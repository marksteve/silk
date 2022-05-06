import { useState, useCallback, useEffect, useReducer } from 'react'
import * as silk from '../wailsjs/go/main/App'

interface Fiber {
  ts: string
  data: any
}

function App() {
  const [status, setStatus] = useState('')
  const [error, setError] = useState(null)

  const [fibers, setFibers] = useReducer(
    (state: Fiber[], action: Fiber[] | Error) =>
      checkError<Fiber[]>(state, action),
    []
  )

  const fetchFibers = useCallback(() => {
    setStatus('Syncing...')
    silk
      .GetFibers()
      .then(setFibers)
      .catch(setError)
      .then(() => setStatus(`Last synced on ${new Date()}`))
  }, [setFibers])

  useEffect(() => {
    setStatus('Connecting to store...')
    // @ts-ignore
    window.runtime.EventsOn('startup', () => {
      setStatus('Connected.')
      fetchFibers()
    })
    // @ts-ignore
    return () => window.runtime.EventsOff('startup')
  }, [])

  return (
    <div className="relative flex h-screen w-screen flex-col overflow-x-hidden">
      <ul className="flex flex-grow flex-col">
        {fibers.map((fiber) => (
          <li key={fiber.ts} className="flex border-b p-5">
            {renderData(fiber.data)}
          </li>
        ))}
      </ul>
      <div className="fixed inset-x-0 bottom-0 flex h-12 items-center justify-between gap-x-5 bg-stone-50 px-5 shadow-inner">
        <div className="flex-grow text-xs">{status}</div>
        <button
          className="rounded-full bg-stone-400 px-2 font-bold text-white"
          onClick={fetchFibers}
        >
          Sync
        </button>
      </div>
      {error ? (
        <div className="absolute inset-0 flex items-center justify-center bg-red-500 p-10 text-white">
          {error}
        </div>
      ) : null}
    </div>
  )
}

function checkError<T>(state: T, action: T | Error) {
  if (action instanceof Error) {
    console.error(action)
    return state
  }
  return action
}

function renderData(data: string) {
  return atob(data)
}

export default App
