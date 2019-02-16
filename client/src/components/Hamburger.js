import React from 'react';

class Hamburger extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
    this.menu = React.createRef();
  }

  handleClick() {
    const classList = this.menu.current.classList; 
    if (classList.contains("on")) {
      classList.remove("on");
      classList.add("off");
    } else {
      classList.remove("off");
      classList.add("on");      
    }
  }

  render() {
    return (
      <div ref={this.menu} className={`hamburger-icon off`} onClick={this.handleClick}>
        <span className='bread-top'>
          <span />
        </span>
        <span className='bread-bottom'>
          <span />
        </span>
      </div>      
    );
  }
}


export default Hamburger;