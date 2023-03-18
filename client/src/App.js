
import Grid from '@mui/material/Unstable_Grid2/Grid2';
import Button from '@mui/material/Button';
import { Box, Divider, Typography, List, ListItem, ListItemButton, Input, Container } from '@mui/material';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';

import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Checkbox from '@mui/material/Checkbox';
import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';



export default function App() {
    const [todos, setTodos] = useState([])
    const [tasks, setTasks] = useState([])
    const [currentTodoId, setCurrentID] = useState(-1)
    const [taskTitleText, setTaskTitleTextInput] = useState('');
    const [todoTitleText, setTodoTitleTextInput] = useState('');
    const [open, setOpen] = React.useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
        setTodoTitleTextInput('');
    };

    useEffect(() => {
        getTodoLists();
    }, [])

    function setCurrentTodo(todoID) {
        setCurrentID(todoID)
        getTasks(todoID)
    }

    function createTodoList(title) {
        var url = "http://localhost:4000/v1/todos/"
        axios.post(url, { "title": title }, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                setTodos([...todos, response.data.new_list]);
            }
        })
    }

    function setTaskCompleted(id) {
        console.log("setting completed " + id)
        var url = "http://localhost:4000/v1/tasks/" + id + "/complete"
        axios.post(url, {}, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                const newTasks = tasks.map(task => {
                    if (task.id === id) {
                        task.completed = true;
                    }
                    return task;
                });
                setTasks(newTasks);
            }
        })
    }

    function deleteTask(id) {
        var url = "http://localhost:4000/v1/tasks/" + id
        axios.delete(url, {}, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                setTasks(tasks.filter(function (x) { return x.id !== id }));
            }
        })
    }
    function deleteTodoList(id) {
        var url = "http://localhost:4000/v1/todos/" + id
        axios.delete(url, {}, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                const newTodos = todos.filter(todo => todo.id !== id);
                setTodos(newTodos);
                if (id === currentTodoId) {
                    if (newTodos.length === 0) {
                        setTasks([]);
                        setCurrentID(-1);
                    } else {
                        setCurrentTodo(newTodos[0].id);
                    }
                }
            }
        })
    }

    function addTaskToCurrentTodo(title) {
        var url = "http://localhost:4000/v1/todos/" + currentTodoId + "/tasks"
        axios.post(url, { "title": title }, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                setTasks(response.data.tasks)
            }
        })
    }
    function getTasks(todo_id) {
        if (todo_id === -1) {
            return
        }
        var url = "http://localhost:4000/v1/todos/" + todo_id + "/tasks"
        axios.get(url, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                setTasks(response.data.tasks)
            }
        })
    }
    function getTodoLists() {
        var url = "http://localhost:4000/v1/todos"
        axios.get(url, {
            mode: 'same-origin',
            responseType: 'json'
        }).then(response => {
            if (response.status === 200) {
                setTodos(response.data.list)
                if (response.data != null &&
                    response.data.list != null &&
                    response.data.list.length > 0 &&
                    currentTodoId === -1) {
                    setCurrentTodo(response.data.list[0].id);
                }
            }
        })
    }

    const getTasksTitle = () => {
        if (currentTodoId === -1) {
            return ""
        }
        return todos.find((x) => x.id === currentTodoId).title;
    }

    const handleTextInputChange = event => {
        setTaskTitleTextInput(event.target.value);
    };

    return <Grid container sx={{ minWidth: 800, minHeight: 500 }}>
        <Grid xs={4}>
            <Box sx={{ bgcolor: '#CCD5AE', width: '100%', height: '100%' }} >
                <Container sx={{ justifyContent: 'center', display: 'flex', height: '10%' }}>
                    <Typography variant='h5'>Todos</Typography>
                </Container>
                <Divider />
                <Container sx={{ height: '80%' }}>
                    <List component="nav" sx={{ width: '100%', justifyContent: 'center' }}>
                        {todos != null && todos.map((todo, i) => (
                            <ListItem secondaryAction={
                                <IconButton edge="end" aria-label="delete"
                                    onClick={function () {
                                        deleteTodoList(todo.id)
                                    }}>
                                    <DeleteIcon />
                                </IconButton>
                            }
                                disablePadding>
                                <ListItemButton
                                    selected={currentTodoId === todo.id}
                                    onClick={function () {
                                        setCurrentTodo(todo.id)
                                    }}>
                                    {todo.title}
                                </ListItemButton>
                            </ListItem>
                        ))}
                    </List>
                </Container>

                <Container sx={{ height: '10%' }}>
                    <Button variant="outlined" onClick={handleClickOpen}>
                        New Todo
                    </Button>
                </Container>
                <Dialog open={open} onClose={handleClose}>
                    <DialogTitle>Subscribe</DialogTitle>
                    <DialogContent>
                        <DialogContentText>
                            Please enter the title of your new Todo List
                        </DialogContentText>
                        <TextField
                            autoFocus
                            margin="dense"
                            value={todoTitleText}
                            onChange={(event) => {
                                setTodoTitleTextInput(event.target.value);
                            }}
                            label="Todo Title"
                            fullWidth
                            variant="standard"
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button onClick={() => {
                            createTodoList(todoTitleText);
                            handleClose();
                        }}>Add</Button>
                    </DialogActions>
                </Dialog>
            </Box>
        </Grid>
        <Grid xs={8}>
            <Box sx={{ bgcolor: '#E9EDC9', width: '100%', height: '100%' }}>
                <Container sx={{ height: '10%' }}>
                    <Typography
                        variant='h4'
                        align="center"
                        sx={{ textDecoration: 'underline' }} >{getTasksTitle()}</Typography>
                </Container>
                <Container sx={{ height: '80%' }}>
                    <Box sx={{ maxHeight: 380, overflow: 'scroll', bgcolor: '#FEFAE0' }}>
                        <List>
                            {tasks != null && tasks.map((task, i) => (
                                <ListItem
                                    secondaryAction={
                                        <IconButton edge="end" aria-label="delete"
                                            onClick={function () {
                                                deleteTask(task.id)
                                            }}>
                                            <DeleteIcon />
                                        </IconButton>
                                    }
                                    disablePadding>
                                    <ListItemButton onClick={
                                        function () {
                                            setTaskCompleted(task.id)
                                        }}
                                    >
                                        <ListItemIcon>
                                            <Checkbox checked={task.completed} />
                                        </ListItemIcon>
                                        <ListItemText>{task.title}</ListItemText>
                                    </ListItemButton>
                                </ListItem>
                            ))}
                        </List>
                    </Box>
                </Container>
                <Container sx={{
                    height: '10%',
                }}>
                    <Input
                        sx={{ width: '85%', align: 'left' }}
                        value={taskTitleText}
                        onChange={handleTextInputChange}>
                    </Input>
                    <Button
                        variant="contained"
                        sx={{
                            float: 'right'
                        }}
                        onClick={
                            function () {
                                addTaskToCurrentTodo(taskTitleText)
                                setTaskTitleTextInput('')
                            }
                        }>Add</Button>
                </Container>
            </Box>
        </Grid>
    </Grid>

}