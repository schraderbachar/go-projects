import { useDisclosure } from '@mantine/hooks';
import { Modal, Button, Group, TextInput, Textarea } from '@mantine/core';
import { ENDPOINT, Todo } from '../App';
import { useForm } from '@mantine/form';
import { KeyedMutator } from 'swr';


function AddTodo({mutate}: {mutate: KeyedMutator<Todo[]>}) {
  const [opened, { open, close }] = useDisclosure(false);


   const form = useForm({
    initialValues: {
      word: '',
      answer: '',
    },
})


  async function createTodo(values: {word: string; answer: string}){
    const updated = await fetch(`${ENDPOINT}/api/todos`, {
        method:'POST',
        headers:{
            "Content-Type":'application/json',
        },
        body: JSON.stringify(values),
    }).then((res) => res.json())
    
    mutate(updated)
    close
    form.reset()
    
}

  return (
    <>
      <Modal opened={opened} onClose={close} title="Authentication">
        <form onSubmit={form.onSubmit(createTodo)}>
            <TextInput required mb={12} label="Word" placeholder='Enter a word you learned' {...form.getInputProps("word")} />
            <Textarea required mb={12} label="body" placeholder='Enter the word in english' {...form.getInputProps("answer")}/>

            <Button type='submit'>Create Word</Button>
        </form>
      </Modal>

      <Group position="center">
        <Button onClick={open}>Add word</Button>
      </Group>
    </>
  );
}
export default AddTodo