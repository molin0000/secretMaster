import styles from './index.css';

function BasicLayout(props) {
  console.log(props);
  let active = styles.menuItem + ' ' + styles.menuItemActive;
  let normal = styles.menuItem;
  let path = props.location.pathname;
  return (
    <div className={styles.normal}>
      <div className={styles.headerRow}>
        <a href="/" className={props.location.pathname === '/' ? active : normal}>概览</a>
        <a href="/login" className={(path === '/config')||(path === '/register')||(path === '/login') ? active : normal }>参数配置</a>
        <a href="/help" className={props.location.pathname === '/help' ? active : normal}>使用手册</a>
      </div>
      {props.children}
    </div>
  );
}

export default BasicLayout;
