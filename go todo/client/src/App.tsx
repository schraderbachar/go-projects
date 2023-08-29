import { Box, Accordion, Button } from '@mantine/core'
import useSWR from "swr"
import './App.css'
import AddTodo from './components/AddTodo'
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';

export interface Todo {
  id: number
  word: string
  answer: string
  done: boolean
}

export const ENDPOINT = 'http://localhost:4000'

const fetcher = (url: string) => fetch(`${ENDPOINT}/${url}`).then(res => res.json())

function App() {
  //fetch data
  const {data,mutate} = useSWR<Todo[]>('api/todos',fetcher)

  //mark todo as done
  async function markDone(id: number){
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
        method: "PATCH",
    }).then(res => res.json())

    mutate(updated)
  }


  return (
    <>
    <Navbar className="bg-body-tertiary">
      <Container>
        <Navbar.Brand className="justify-content-center">GoLang Practice</Navbar.Brand>
      </Container>
    </Navbar>
   <Box sx={() => ({
    padding: "2rem",
    width: "100vw",
    maxWidth: "40rem",
    margin: "0 auto",
    color: "#fff"
   })}>
    <Accordion variant="separated" styles={{
        item: {
          // styles added to all items
          backgroundColor: '#fff',
          border: '2px solid black',
          // styles added to expanded item
          '&[data-active]': {
            backgroundColor: '#ccc',
            color: '#000'
          },
        },

        chevron: {
          // styles added to chevron when it should rotate
          '&[data-rotate]': {
            transform: 'rotate(-90deg)',
          },
        },
      }}
    >
    {data?.map(todo => {
      return (
        <>{!todo.done && <Accordion.Item key={todo.id} value={todo.word}>
            <Accordion.Control>{todo.word}</Accordion.Control>
            <Accordion.Panel>{todo.answer} {todo.done}
              <Button color='green' onClick={() => markDone(todo.id)}>Learned</Button>
            </Accordion.Panel>
          </Accordion.Item>
          }
      </>)
      })} 
      </Accordion>
      <p className='m-5'></p>
    <AddTodo mutate={mutate} />
   </Box>
   </>
  )
}

export default App
