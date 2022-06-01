func s:create_popup()
  let id = popup_menu('My popup', #{
	\ pos: 'topright',
	\ padding: [2, 2],
	\ border: [],
	\ borderchars: ['─', '│', '─', '│', '╭', '╮', '╯', '╰'],
	\ title: 'Hello World'
	\ })
endfunc

nnoremap ,,a :call<SID>create_popup()<CR>
