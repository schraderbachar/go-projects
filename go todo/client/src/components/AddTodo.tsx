import { useDisclosure } from '@mantine/hooks';
import { Modal, Button, Group, TextInput, Textarea } from '@mantine/core';
import { ENDPOINT, Todo } from '../App';
import { useForm } from '@mantine/form';
import { KeyedMutator } from 'swr';


function AddTodo({mutate}: {mutate: KeyedMutator<Todo[]>}) {
  const [opened, { open, close }] = useDisclosure(false);


   const form = useForm({
    initialValues: {
      title: '',
      body: '',
    },
})


  async function createTodo(values: {title: string; body: string}){
    const updated = await fetch(`${ENDPOINT}/api/todos`, {
        method:'POST',
        headers:{
            "Content-Type":'application/json',
        },
        body: JSON.stringify(values),
    }).then((res) => res.json())
    
    mutate(updated)
    form.reset()
    close
}

  return (
    <>
      <Modal opened={opened} onClose={close} title="Authentication">
        <form onSubmit={form.onSubmit(createTodo)}>
            <TextInput required mb={12} label="todo" placeholder='what do you want to do?' {...form.getInputProps("title")} />
            <Textarea required mb={12} label="body" placeholder='Tell me more' {...form.getInputProps("body")}/>

            <Button type='submit'>Create todo</Button>
        </form>
      </Modal>

      <Group position="center">
        <Button onClick={open}>Open modal</Button>
      </Group>
    </>
  );
}
export default AddTodo