import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Checkbox from '@mui/material/Checkbox';

import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';
import axios from 'axios';
import { useState, useEffect } from 'react';

import { Box, Button, Typography, List, ListItem, ListItemButton, Input, Container } from '@mui/material';

export default function TasksComponent(currentTodoId) {
    const [tasks, setTasks] = useState([])
    const [taskTitleText, setTaskTitleTextInput] = useState('');

    useEffect(() => {
        getTasks();
    }, [])

    function getTasks() {
        if (currentTodoId === -1) {
            return
        }
        var url = "http://localhost:4000/v1/todos/" + currentTodoId + "/tasks"
        axios.get(url, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                setTasks(response.data.tasks)
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

    function setTaskCompleted(id) {
        console.log("setting completed " + id)
        var url = "http://localhost:4000/v1/tasks/" + id + "/complete"
        axios.post(url, {}, {
            mode: 'same-origin',
            responseType: 'json',
        }).then(response => {
            if (response.status === 200) {
                const index = tasks.indexOf(id);
                tasks[index].completed = true;
                setTasks(tasks);
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

    const handleTextInputChange = event => {
        setTaskTitleTextInput(event.target.value);
    };

    return <> <Container sx={{ height: '90%' }}>
        <Typography variant='h4'>Tasks</Typography>
        <Box>
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
        <Container sx={{ height: '10%', spacing: 50 }}>
            <Input
                sx={{ width: '80%' }}
                value={taskTitleText}
                onChange={handleTextInputChange}>
            </Input>
            <Button
                variant="contained"
                sx={{ alignSelf: 'end' }}
                onClick={
                    function () {
                        addTaskToCurrentTodo(taskTitleText)
                    }
                }>Add</Button>
        </Container></>
}