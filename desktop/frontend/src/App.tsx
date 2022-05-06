import {
  HTMLAttributes,
  useState,
  useCallback,
  useEffect,
  useReducer,
} from 'react'
import { format } from 'timeago.js'
import * as silk from '../wailsjs/go/main/App'

interface Fiber {
  ts: string
  data: any
  mimetype: string
}

const Button = (props: HTMLAttributes<HTMLButtonElement>) => (
  <button
    {...props}
    className={`rounded-full bg-stone-400 px-3 text-sm font-bold leading-loose text-white hover:bg-stone-600 ${props.className}`}
  />
)

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
      .then((resp) => setFibers(resp ?? []))
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
    window.addEventListener('paste', (e: any) => {
      setStatus('Pasting...')
      const data = btoa(e.clipboardData.getData('text'))
      // @ts-ignore
      silk.Weave(data).then(fetchFibers)
    })
    return () => {
      // @ts-ignore
      window.runtime.EventsOff('startup')
    }
  }, [])

  return (
    <div className="relative flex h-screen w-screen flex-col overflow-x-hidden">
      <ul className="overflow-y-auto" style={{ height: 'calc(100vh - 3rem)' }}>
        {fibers.map((fiber) => (
          <li
            key={fiber.ts}
            className="group flex max-h-32 items-start justify-between overflow-hidden border-b p-5 hover:bg-stone-50"
          >
            <div className="flex w-11/12 flex-col">
              <div className="text-xs text-stone-500">{format(fiber.ts)}</div>
              {renderData(fiber.data)}
            </div>
            <Button
              className="invisible group-hover:visible"
              onClick={() => copyToClipboard(fiber)}
            >
              Copy
            </Button>
          </li>
        ))}
      </ul>
      <div className="flex h-12 items-center justify-between gap-x-5 bg-stone-50 px-5 shadow-inner">
        <div className="w-11/12 text-xs text-stone-500">{status}</div>
        <Button onClick={fetchFibers}>Sync</Button>
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
  return <div className="whitespace-pre">{atob(data)}</div>
}

function copyToClipboard(fiber: Fiber) {
  navigator.clipboard.writeText(atob(fiber.data))
}

export default App
