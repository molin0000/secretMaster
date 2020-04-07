import styles from './index.css';

function BasicLayout(props) {
  console.log(props);
  let active = styles.menuItem + ' ' + styles.menuItemActive;
  let normal = styles.menuItem;
  let path = props.location.pathname;
  document.title = '⭐️序列战争⭐️';
  return (
    <div className={styles.normal}>
      <div className={styles.headerRow}>
        <a href="/" className={props.location.pathname === '/' ? active : normal}>游玩</a>
        <a href="/config" className={(path === '/config')||(path === '/register')||(path === '/login') ? active : normal }>设置</a>
        <a href="/help" className={props.location.pathname === '/help' ? active : normal}>帮助</a>
        <a href="/about" className={props.location.pathname === '/about' ? active : normal}>关于</a>
      </div>
      {props.children}
    </div>
  );
}

export default BasicLayout;
