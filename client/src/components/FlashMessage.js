import React from 'react';

class FlashMessage extends React.Component() {
  constructor(props) {
    super(props);
    this.state = {
      visible: true
    }
  }

  componentDidMount() {
    let self = this;
    setTimeout(() => {
      self.setState({
        visible: false
      })
    }, 5000)
  }

  render() {
    if (!this.state.visible) { return null; }
    return (
      <div>I AM MESSAGE</div>
    );
  }
}

export default FlashMessage;