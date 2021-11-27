import { Component, Host, h, State, Listen } from '@stencil/core';

interface TodoItem {
  ID?: number;
  text: string;
  checked: boolean;
}

@Component({
  tag: 'todo-list',
  styleUrl: 'todo-list.css',
  shadow: true,
})
export class TodoList {
  @State() list: TodoItem[]

  @Listen('onTodoInputSubmit')
  async todoInputSubmiHandler(e: CustomEvent) {
    const res = await fetch('http://localhost:8080/todos', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        text: e.detail
      })
    });
    this.list = [...this.list, await res.json()];
  }

  @Listen('onTodoItemChecked')
  async todoItemCheckedHandler(e: CustomEvent) {
    const list = [...this.list];
    const item = list[e.detail];
    list[e.detail] = Object.assign({}, item, { checked: !item.checked });
    this.list = list;
    await fetch(`http://localhost:8080/todos/${item.ID}/check`, {method: 'POST'});
  }

  @Listen('onTodoItemRemove')
  async todoItemRemoveHandler(e: CustomEvent) {
    const item = this.list[e.detail];
    await fetch(`http://localhost:8080/todos/${item.ID}`, {method: 'DELETE'});
    this.list = [...this.list.slice(0, e.detail), ...this.list.slice(e.detail + 1)];
  }

  async componentWillLoad() {
    const res = await fetch('http://localhost:8080/todos');
    this.list = await res.json();
  }

  render() {
    return (
      <Host>
        <div>
          <h1>Todos app</h1>
          <section>
            <todo-input></todo-input>
            <ul id="list-container">
              {this.list.map((item, index) => (
                <todo-item
                  checked={item.checked}
                  text={item.text}
                  index={index}
                />
              ))}
            </ul>
          </section>
        </div>
      </Host>
    );
  }
}
