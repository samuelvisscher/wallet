let SessionLoad = 1
if &cp | set nocp | endif
let s:cpo_save=&cpo
set cpo&vim
inoremap <expr> <Down> pumvisible() ? "\" : "\<Down>"
inoremap <expr> <S-Tab> pumvisible() ? "\" : "\<S-Tab>"
inoremap <expr> <Up> pumvisible() ? "\" : "\<Up>"
imap <Nul> <C-Space>
inoremap <silent> <SNR>33_AutoPairsReturn =AutoPairsReturn()
nnoremap <silent>  :TmuxNavigateLeft
nnoremap <silent> <NL> :TmuxNavigateDown
nnoremap <silent>  :TmuxNavigateUp
nnoremap <silent>  :TmuxNavigateRight
map  <Plug>(ctrlp)
nnoremap <silent>  :TmuxNavigatePrevious
nmap <PasteEnd> <Nop>
nmap <PasteStart> <Nop>
noremap \<F2> :wq:Gpush
nnoremap \d :YcmShowDetailedDiagnostic
map \	 :exec &conceallevel ? "set conceallevel=0" : "set conceallevel=1"
map \y :CtrlPBuffer
map \t 
map \- gg=G``zz
noremap \  :Autoformat
map \cs :Tabularize /:\zs
map \es :Tabularize /=\zs
map \c :Tabularize /:
map \e :Tabularize /=
map \' :NERDTreeToggle
noremap <silent> \\ :call ToggleRelativeOn()
map \l 
map \b :b
map \p :bp
map \n :bn
map \i "*p
map \x :q
map \s :w
noremap \<F12> :tabNext
noremap \<F10> :wq:Gpush
noremap \<F9> :Git add %:Gcommiti
noremap \<F8> :Git diff %
noremap \<F7> :Gpull
noremap \<F6> :Gstatus
noremap \<F5> :wq:Gpush
noremap \<F4> :Gcommiti
noremap \<F3> :Git add %
map \ws :ChooseWin
map \ct :w:!cucumber
map \cl :w:exe "!cucumber %" . ":" . line(".")
map \cf :w:!cucumber %
map \rt :w:!rspec --format nested
map \rl :w:exe "!rspec %" . ":" . line(".")
map \rf :w:!rspec % --format nested
map \` v
map \; s
map \" :e ~/.vimrc
map \m :PluginInstall
map \v :source ~/.vimrc
map \= k " Move to left window
map \ j " Move to left window
map \[ h " Move to left window
map \] l " Move to right window
map \w[ = " equalize all windows
map \w] _ " maximize height
map \w+ + " increment height
map \w- - " decrement height
map \yd :bufdo bd
map \yt :ls
map \. :Sexplore
nmap \! %!sudo tee > /dev/null %
nmap <silent> b <Plug>CamelCaseMotion_b
xmap <silent> b <Plug>CamelCaseMotion_b
omap <silent> b <Plug>CamelCaseMotion_b
nmap <silent> e <Plug>CamelCaseMotion_e
xmap <silent> e <Plug>CamelCaseMotion_e
omap <silent> e <Plug>CamelCaseMotion_e
vmap gx <Plug>NetrwBrowseXVis
nmap gx <Plug>NetrwBrowseX
nmap <silent> w <Plug>CamelCaseMotion_w
xmap <silent> w <Plug>CamelCaseMotion_w
omap <silent> w <Plug>CamelCaseMotion_w
vnoremap <silent> <Plug>NetrwBrowseXVis :call netrw#BrowseXVis()
nnoremap <silent> <Plug>NetrwBrowseX :call netrw#BrowseX(expand((exists("g:netrw_gx")? g:netrw_gx : '<cfile>')),netrw#CheckIfRemote())
nnoremap <silent> <Plug>(JavaComplete-Imports-SortImports) :call javacomplete#imports#SortImports()
nnoremap <silent> <Plug>(JavaComplete-Generate-ClassInFile) :call javacomplete#newclass#CreateInFile()
nnoremap <silent> <Plug>(JavaComplete-Generate-NewClass) :call javacomplete#newclass#CreateClass()
nnoremap <silent> <Plug>(JavaComplete-Generate-DefaultConstructor) :call javacomplete#generators#GenerateConstructor(1)
nnoremap <silent> <Plug>(JavaComplete-Generate-Constructor) :call javacomplete#generators#GenerateConstructor(0)
nnoremap <silent> <Plug>(JavaComplete-Generate-EqualsAndHashCode) :call javacomplete#generators#GenerateEqualsAndHashCode()
nnoremap <silent> <Plug>(JavaComplete-Generate-ToString) :call javacomplete#generators#GenerateToString()
vnoremap <silent> <Plug>(JavaComplete-Generate-AccessorSetterGetter) :call javacomplete#generators#Accessor('sg')
vnoremap <silent> <Plug>(JavaComplete-Generate-AccessorGetter) :call javacomplete#generators#Accessor('g')
vnoremap <silent> <Plug>(JavaComplete-Generate-AccessorSetter) :call javacomplete#generators#Accessor('s')
nnoremap <silent> <Plug>(JavaComplete-Generate-AccessorSetterGetter) :call javacomplete#generators#Accessor('sg')
nnoremap <silent> <Plug>(JavaComplete-Generate-AccessorGetter) :call javacomplete#generators#Accessor('g')
nnoremap <silent> <Plug>(JavaComplete-Generate-AccessorSetter) :call javacomplete#generators#Accessor('s')
nnoremap <silent> <Plug>(JavaComplete-Generate-Accessors) :call javacomplete#generators#Accessors()
nnoremap <silent> <Plug>(JavaComplete-Generate-AbstractMethods) :call javacomplete#generators#AbstractDeclaration()
nnoremap <silent> <Plug>(JavaComplete-Imports-AddSmart) :call javacomplete#imports#Add(1)
nnoremap <silent> <Plug>(JavaComplete-Imports-Add) :call javacomplete#imports#Add()
nnoremap <silent> <Plug>(JavaComplete-Imports-RemoveUnused) :call javacomplete#imports#RemoveUnused()
nnoremap <silent> <Plug>(JavaComplete-Imports-AddMissing) :call javacomplete#imports#AddMissing()
nnoremap <SNR>82_: :=v:count ? v:count : ''
nnoremap <silent> <Plug>(choosewin) :call choosewin#start(range(1, winnr('$')))
nnoremap <silent> <Plug>(ctrlp) :CtrlP
vnoremap <silent> <Plug>CamelCaseMotion_ige :call camelcasemotion#InnerMotion('ge',v:count1)
vnoremap <silent> <Plug>CamelCaseMotion_ie :call camelcasemotion#InnerMotion('e',v:count1)
vnoremap <silent> <Plug>CamelCaseMotion_ib :call camelcasemotion#InnerMotion('b',v:count1)
vnoremap <silent> <Plug>CamelCaseMotion_iw :call camelcasemotion#InnerMotion('w',v:count1)
onoremap <silent> <Plug>CamelCaseMotion_ige :call camelcasemotion#InnerMotion('ge',v:count1)
onoremap <silent> <Plug>CamelCaseMotion_ie :call camelcasemotion#InnerMotion('e',v:count1)
onoremap <silent> <Plug>CamelCaseMotion_ib :call camelcasemotion#InnerMotion('b',v:count1)
onoremap <silent> <Plug>CamelCaseMotion_iw :call camelcasemotion#InnerMotion('w',v:count1)
vnoremap <silent> <Plug>CamelCaseMotion_ge :call camelcasemotion#Motion('ge',v:count1,'v')
vnoremap <silent> <Plug>CamelCaseMotion_e :call camelcasemotion#Motion('e',v:count1,'v')
vnoremap <silent> <Plug>CamelCaseMotion_b :call camelcasemotion#Motion('b',v:count1,'v')
vnoremap <silent> <Plug>CamelCaseMotion_w :call camelcasemotion#Motion('w',v:count1,'v')
onoremap <silent> <Plug>CamelCaseMotion_ge :call camelcasemotion#Motion('ge',v:count1,'o')
onoremap <silent> <Plug>CamelCaseMotion_e :call camelcasemotion#Motion('e',v:count1,'o')
onoremap <silent> <Plug>CamelCaseMotion_b :call camelcasemotion#Motion('b',v:count1,'o')
onoremap <silent> <Plug>CamelCaseMotion_w :call camelcasemotion#Motion('w',v:count1,'o')
nnoremap <silent> <Plug>CamelCaseMotion_ge :call camelcasemotion#Motion('ge',v:count1,'n')
nnoremap <silent> <Plug>CamelCaseMotion_e :call camelcasemotion#Motion('e',v:count1,'n')
nnoremap <silent> <Plug>CamelCaseMotion_b :call camelcasemotion#Motion('b',v:count1,'n')
nnoremap <silent> <Plug>CamelCaseMotion_w :call camelcasemotion#Motion('w',v:count1,'n')
nmap <F8> :TagbarToggle
nnoremap <F2> :set invpaste paste?
inoremap <expr> 	 pumvisible() ? "\" : "\	"
cmap w!! :w !sudo tee >/dev/null %
let &cpo=s:cpo_save
unlet s:cpo_save
set autoindent
set autoread
set backspace=indent,eol,start
set clipboard=autoselect,exclude:cons\\|linux,unnamed
set completefunc=youcompleteme#CompleteFunc
set completeopt=preview,menuone
set expandtab
set fileencodings=ucs-bom,utf-8,default,latin1
set guioptions=aegimrLt
set helplang=en
set hidden
set history=100
set hlsearch
set ignorecase
set incsearch
set laststatus=2
set lazyredraw
set listchars=tab:▸\ ,eol:¬,trail:·
set mouse=a
set pastetoggle=<F2>
set printoptions=paper:letter
set ruler
set runtimepath=~/.vim,~/.vim/bundle/Vundle.vim,~/.vim/bundle/vim-colorschemes,~/.vim/bundle/ctrlp.vim,~/.vim/bundle/editorconfig-vim,~/.vim/bundle/vim-airline,~/.vim/bundle/vim-airline-themes,~/.vim/bundle/nerdtree,~/.vim/bundle/vim-autoformat,~/.vim/bundle/tabular,~/.vim/bundle/vim-choosewin,~/.vim/bundle/vimproc.vim,~/.vim/bundle/vim-rspec,~/.vim/bundle/YouCompleteMe,~/.vim/bundle/tagbar,~/.vim/bundle/ctags,~/.vim/bundle/tagbar-phpctags.vim,~/.vim/bundle/nerdtree-git-plugin,~/.vim/bundle/vim-fugitive,~/.vim/bundle/auto-pairs,~/.vim/bundle/CamelCaseMotion,~/.vim/bundle/rainbow,~/.vim/bundle/angular-cli.vim,~/.vim/bundle/vim-angular-template,~/.vim/bundle/vim-javacomplete2,~/.vim/bundle/vim-js-pretty-template,~/.vim/bundle/vim-javascript,~/.vim/bundle/vim-jsx,~/.vim/bundle/vim-json,~/.vim/bundle/typescript-vim,~/.vim/bundle/vim-tmux-navigator,~/.vim/bundle/vim-vue,~/.vim/bundle/CamelCaseMotion,~/.vim/bundle/Vundle.vim,~/.vim/bundle/YouCompleteMe,~/.vim/bundle/angular-cli.vim,~/.vim/bundle/auto-pairs,~/.vim/bundle/ctags,~/.vim/bundle/ctrlp.vim,~/.vim/bundle/editorconfig-vim,~/.vim/bundle/nerdtree,~/.vim/bundle/nerdtree-git-plugin,~/.vim/bundle/pathogen,~/.vim/bundle/rainbow,~/.vim/bundle/syntastic,~/.vim/bundle/tabular,~/.vim/bundle/tagbar,~/.vim/bundle/tagbar-phpctags.vim,~/.vim/bundle/typescript-vim,~/.vim/bundle/vim-airline,~/.vim/bundle/vim-airline-themes,~/.vim/bundle/vim-angular-template,~/.vim/bundle/vim-autoformat,~/.vim/bundle/vim-choosewin,~/.vim/bundle/vim-colorschemes,~/.vim/bundle/vim-fugitive,~/.vim/bundle/vim-gradle,~/.vim/bundle/vim-javacomplete2,~/.vim/bundle/vim-javascript,~/.vim/bundle/vim-js-pretty-template,~/.vim/bundle/vim-json,~/.vim/bundle/vim-jsx,~/.vim/bundle/vim-rspec,~/.vim/bundle/vim-tmux-navigator,~/.vim/bundle/vim-vue,~/.vim/bundle/vimproc.vim,/var/lib/vim/addons,/usr/share/vim/vimfiles,/usr/share/vim/vim80,/usr/share/vim/vimfiles/after,/var/lib/vim/addons/after,~/.vim/bundle/vim-jsx/after,~/.vim/bundle/vim-javascript/after,~/.vim/bundle/vim-gradle/after,~/.vim/bundle/vim-angular-template/after,~/.vim/bundle/tabular/after,~/.vim/after,~/.vim/bundle/Vundle.vim/after,~/.vim/bundle/vim-colorschemes/after,~/.vim/bundle/ctrlp.vim/after,~/.vim/bundle/editorconfig-vim/after,~/.vim/bundle/vim-airline/after,~/.vim/bundle/vim-airline-themes/after,~/.vim/bundle/nerdtree/after,~/.vim/bundle/vim-autoformat/after,~/.vim/bundle/vim-choosewin/after,~/.vim/bundle/vimproc.vim/after,~/.vim/bundle/vim-rspec/after,~/.vim/bundle/YouCompleteMe/after,~/.vim/bundle/tagbar/after,~/.vim/bundle/ctags/after,~/.vim/bundle/tagbar-phpctags.vim/after,~/.vim/bundle/nerdtree-git-plugin/after,~/.vim/bundle/vim-fugitive/after,~/.vim/bundle/auto-pairs/after,~/.vim/bundle/CamelCaseMotion/after,~/.vim/bundle/rainbow/after,~/.vim/bundle/angular-cli.vim/after,~/.vim/bundle/vim-javacomplete2/after,~/.vim/bundle/vim-js-pretty-template/after,~/.vim/bundle/vim-json/after,~/.vim/bundle/typescript-vim/after,~/.vim/bundle/vim-tmux-navigator/after,~/.vim/bundle/vim-vue/after,~/.vim/bundle/Vundle.vim/after,~/.vim/bundle/vim-colorschemes/after,~/.vim/bundle/ctrlp.vim/after,~/.vim/bundle/editorconfig-vim/after,~/.vim/bundle/vim-airline/after,~/.vim/bundle/vim-airline-themes/after,~/.vim/bundle/nerdtree/after,~/.vim/bundle/vim-autoformat/after,~/.vim/bundle/tabular/after,~/.vim/bundle/vim-choosewin/after,~/.vim/bundle/vimproc.vim/after,~/.vim/bundle/vim-rspec/after,~/.vim/bundle/YouCompleteMe/after,~/.vim/bundle/tagbar/after,~/.vim/bundle/ctags/after,~/.vim/bundle/tagbar-phpctags.vim/after,~/.vim/bundle/nerdtree-git-plugin/after,~/.vim/bundle/vim-fugitive/after,~/.vim/bundle/auto-pairs/after,~/.vim/bundle/CamelCaseMotion/after,~/.vim/bundle/rainbow/after,~/.vim/bundle/angular-cli.vim/after,~/.vim/bundle/vim-angular-template/after,~/.vim/bundle/vim-javacomplete2/after,~/.vim/bundle/vim-js-pretty-template/after,~/.vim/bundle/vim-javascript/after,~/.vim/bundle/vim-jsx/after,~/.vim/bundle/vim-json/after,~/.vim/bundle/typescript-vim/after,~/.vim/bundle/vim-tmux-navigator/after,~/.vim/bundle/vim-vue/after
set shiftwidth=2
set shortmess=filnxtToOI
set showcmd
set showmatch
set showtabline=2
set smartcase
set splitbelow
set splitright
set statusline=%f\ %=L:%l/%L\ %c\ (%p%%)
set suffixes=.bak,~,.swp,.o,.info,.aux,.log,.dvi,.bbl,.blg,.brf,.cb,.ind,.idx,.ilg,.inx,.out,.toc
set noswapfile
set tabline=%!airline#extensions#tabline#get()
set tabstop=2
set ttimeoutlen=100
set visualbell
set wildignore=*/.git/*,*/.hg/*,*/.svn/*.,*/.DS_Store
set wildmenu
set nowritebackup
let s:so_save = &so | let s:siso_save = &siso | set so=0 siso=0
let v:this_session=expand("<sfile>:p")
silent only
cd ~/.go/src/wallet
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
badd +1 src/iko/blockchain.go
badd +38 src/iko/chain.go
badd +1 src/iko/kitty.go
badd +5 src/iko/kitty_test.go
badd +79 src/iko/state.go
badd +1 src/iko/state_test.go
badd +136 src/iko/transaction.go
badd +38 src/iko/transaction_test.go
argglobal
silent! argdel *
argadd src/iko/blockchain.go
argadd src/iko/transaction.go
argadd src/iko/state_test.go
argadd src/iko/state.go
argadd src/iko/kitty_test.go
argadd src/iko/kitty.go
argadd src/iko/chain.go
edit src/iko/transaction.go
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd t
set winheight=1 winwidth=1
exe 'vert 1resize ' . ((&columns * 118 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 119 + 119) / 238)
argglobal
edit src/iko/transaction.go
let s:cpo_save=&cpo
set cpo&vim
inoremap <buffer> <silent> <BS> =AutoPairsDelete()
inoremap <buffer> <silent> § =AutoPairsMoveCharacter('''')
inoremap <buffer> <silent> ¢ =AutoPairsMoveCharacter('"')
inoremap <buffer> <silent> © =AutoPairsMoveCharacter(')')
inoremap <buffer> <silent> ¨ =AutoPairsMoveCharacter('(')
inoremap <buffer> <silent> î :call AutoPairsJump()a
inoremap <buffer> <silent> <expr> ð AutoPairsToggle()
inoremap <buffer> <silent> â =AutoPairsBackInsert()
inoremap <buffer> <silent> å =AutoPairsFastWrap()
inoremap <buffer> <silent> ý =AutoPairsMoveCharacter('}')
inoremap <buffer> <silent> û =AutoPairsMoveCharacter('{')
inoremap <buffer> <silent> Ý =AutoPairsMoveCharacter(']')
inoremap <buffer> <silent> Û =AutoPairsMoveCharacter('[')
inoremap <buffer> <silent>  =AutoPairsDelete()
inoremap <buffer> <silent>   =AutoPairsSpace()
inoremap <buffer> <silent> " =AutoPairsInsert('"')
inoremap <buffer> <silent> ' =AutoPairsInsert('''')
inoremap <buffer> <silent> ( =AutoPairsInsert('(')
inoremap <buffer> <silent> ) =AutoPairsInsert(')')
noremap <buffer> <silent> î :call AutoPairsJump()
noremap <buffer> <silent> ð :call AutoPairsToggle()
inoremap <buffer> <silent> [ =AutoPairsInsert('[')
inoremap <buffer> <silent> ] =AutoPairsInsert(']')
inoremap <buffer> <silent> ` =AutoPairsInsert('`')
inoremap <buffer> <silent> { =AutoPairsInsert('{')
inoremap <buffer> <silent> } =AutoPairsInsert('}')
let &cpo=s:cpo_save
unlet s:cpo_save
setlocal keymap=
setlocal noarabic
setlocal autoindent
setlocal backupcopy=
setlocal balloonexpr=
setlocal nobinary
setlocal nobreakindent
setlocal breakindentopt=
setlocal bufhidden=
setlocal buflisted
setlocal buftype=
setlocal nocindent
setlocal cinkeys=0{,0},0),:,0#,!^F,o,O,e
setlocal cinoptions=
setlocal cinwords=if,else,while,do,for,switch
set colorcolumn=80
setlocal colorcolumn=80
setlocal comments=s1:/*,mb:*,ex:*/,://
setlocal commentstring=//\ %s
setlocal complete=.,w,b,u,t,i
setlocal concealcursor=
setlocal conceallevel=0
setlocal completefunc=youcompleteme#CompleteFunc
setlocal nocopyindent
setlocal cryptmethod=
setlocal nocursorbind
setlocal nocursorcolumn
set cursorline
setlocal cursorline
setlocal define=
setlocal dictionary=
setlocal nodiff
setlocal equalprg=
setlocal errorformat=
setlocal noexpandtab
if &filetype != 'go'
setlocal filetype=go
endif
setlocal nofixendofline
setlocal foldcolumn=0
setlocal foldenable
setlocal foldexpr=0
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldmarker={{{,}}}
setlocal foldmethod=marker
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldtext=foldtext()
setlocal formatexpr=
setlocal formatoptions=cq
setlocal formatlistpat=^\\s*\\d\\+[\\]:.)}\\t\ ]\\s*
setlocal formatprg=
setlocal grepprg=
setlocal iminsert=2
setlocal imsearch=2
setlocal include=
setlocal includeexpr=
setlocal indentexpr=GoIndent(v:lnum)
setlocal indentkeys=0{,0},:,0#,!^F,o,O,e,<:>,0=},0=)
setlocal noinfercase
setlocal iskeyword=@,48-57,_,192-255
setlocal keywordprg=
setlocal nolinebreak
setlocal nolisp
setlocal lispwords=
set list
setlocal list
setlocal makeprg=
setlocal matchpairs=(:),{:},[:]
setlocal modeline
setlocal modifiable
setlocal nrformats=bin,octal,hex
set number
setlocal number
setlocal numberwidth=4
setlocal omnifunc=
setlocal path=
setlocal nopreserveindent
setlocal nopreviewwindow
setlocal quoteescape=\\
setlocal noreadonly
set relativenumber
setlocal relativenumber
setlocal norightleft
setlocal rightleftcmd=search
setlocal noscrollbind
setlocal shiftwidth=4
setlocal noshortname
setlocal signcolumn=auto
setlocal nosmartindent
setlocal softtabstop=4
setlocal nospell
setlocal spellcapcheck=[.?!]\\_[\\])'\"\	\ ]\\+
setlocal spellfile=
setlocal spelllang=en
setlocal statusline=%!airline#statusline(1)
setlocal suffixesadd=
setlocal noswapfile
setlocal synmaxcol=3000
if &syntax != 'go'
setlocal syntax=go
endif
setlocal tabstop=4
setlocal tagcase=
setlocal tags=
setlocal textwidth=0
setlocal thesaurus=
setlocal noundofile
setlocal undolevels=-123456
setlocal nowinfixheight
setlocal nowinfixwidth
set nowrap
setlocal nowrap
setlocal wrapmargin=0
let s:l = 52 - ((46 * winheight(0) + 25) / 50)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
52
normal! 0
wincmd w
argglobal
edit src/iko/transaction_test.go
let s:cpo_save=&cpo
set cpo&vim
inoremap <buffer> <silent> <BS> =AutoPairsDelete()
inoremap <buffer> <silent> § =AutoPairsMoveCharacter('''')
inoremap <buffer> <silent> ¢ =AutoPairsMoveCharacter('"')
inoremap <buffer> <silent> © =AutoPairsMoveCharacter(')')
inoremap <buffer> <silent> ¨ =AutoPairsMoveCharacter('(')
inoremap <buffer> <silent> î :call AutoPairsJump()a
inoremap <buffer> <silent> <expr> ð AutoPairsToggle()
inoremap <buffer> <silent> â =AutoPairsBackInsert()
inoremap <buffer> <silent> å =AutoPairsFastWrap()
inoremap <buffer> <silent> ý =AutoPairsMoveCharacter('}')
inoremap <buffer> <silent> û =AutoPairsMoveCharacter('{')
inoremap <buffer> <silent> Ý =AutoPairsMoveCharacter(']')
inoremap <buffer> <silent> Û =AutoPairsMoveCharacter('[')
inoremap <buffer> <silent>  =AutoPairsDelete()
inoremap <buffer> <silent>   =AutoPairsSpace()
inoremap <buffer> <silent> " =AutoPairsInsert('"')
inoremap <buffer> <silent> ' =AutoPairsInsert('''')
inoremap <buffer> <silent> ( =AutoPairsInsert('(')
inoremap <buffer> <silent> ) =AutoPairsInsert(')')
noremap <buffer> <silent> î :call AutoPairsJump()
noremap <buffer> <silent> ð :call AutoPairsToggle()
inoremap <buffer> <silent> [ =AutoPairsInsert('[')
inoremap <buffer> <silent> ] =AutoPairsInsert(']')
inoremap <buffer> <silent> ` =AutoPairsInsert('`')
inoremap <buffer> <silent> { =AutoPairsInsert('{')
inoremap <buffer> <silent> } =AutoPairsInsert('}')
let &cpo=s:cpo_save
unlet s:cpo_save
setlocal keymap=
setlocal noarabic
setlocal autoindent
setlocal backupcopy=
setlocal balloonexpr=
setlocal nobinary
setlocal nobreakindent
setlocal breakindentopt=
setlocal bufhidden=
setlocal buflisted
setlocal buftype=
setlocal nocindent
setlocal cinkeys=0{,0},0),:,0#,!^F,o,O,e
setlocal cinoptions=
setlocal cinwords=if,else,while,do,for,switch
set colorcolumn=80
setlocal colorcolumn=80
setlocal comments=s1:/*,mb:*,ex:*/,://
setlocal commentstring=//\ %s
setlocal complete=.,w,b,u,t,i
setlocal concealcursor=
setlocal conceallevel=0
setlocal completefunc=youcompleteme#CompleteFunc
setlocal nocopyindent
setlocal cryptmethod=
setlocal nocursorbind
setlocal nocursorcolumn
set cursorline
setlocal cursorline
setlocal define=
setlocal dictionary=
setlocal nodiff
setlocal equalprg=
setlocal errorformat=
setlocal expandtab
if &filetype != 'go'
setlocal filetype=go
endif
setlocal nofixendofline
setlocal foldcolumn=0
setlocal foldenable
setlocal foldexpr=0
setlocal foldignore=#
setlocal foldlevel=0
setlocal foldmarker={{{,}}}
setlocal foldmethod=manual
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldtext=foldtext()
setlocal formatexpr=
setlocal formatoptions=tcq
setlocal formatlistpat=^\\s*\\d\\+[\\]:.)}\\t\ ]\\s*
setlocal formatprg=
setlocal grepprg=
setlocal iminsert=2
setlocal imsearch=2
setlocal include=
setlocal includeexpr=
setlocal indentexpr=GoIndent(v:lnum)
setlocal indentkeys=0{,0},:,0#,!^F,o,O,e,<:>,0=},0=)
setlocal noinfercase
setlocal iskeyword=@,48-57,_,192-255
setlocal keywordprg=
setlocal nolinebreak
setlocal nolisp
setlocal lispwords=
set list
setlocal list
setlocal makeprg=
setlocal matchpairs=(:),{:},[:]
setlocal modeline
setlocal modifiable
setlocal nrformats=bin,octal,hex
set number
setlocal number
setlocal numberwidth=4
setlocal omnifunc=
setlocal path=
setlocal nopreserveindent
setlocal previewwindow
setlocal quoteescape=\\
setlocal noreadonly
setlocal relativenumber
setlocal norightleft
setlocal rightleftcmd=search
setlocal noscrollbind
setlocal shiftwidth=2
setlocal noshortname
setlocal signcolumn=auto
setlocal nosmartindent
setlocal softtabstop=4
setlocal nospell
setlocal spellcapcheck=[.?!]\\_[\\])'\"\	\ ]\\+
setlocal spellfile=
setlocal spelllang=en
setlocal statusline=%!airline#statusline(2)
setlocal suffixesadd=
setlocal noswapfile
setlocal synmaxcol=3000
if &syntax != 'go'
setlocal syntax=go
endif
setlocal tabstop=2
setlocal tagcase=
setlocal tags=
setlocal textwidth=0
setlocal thesaurus=
setlocal noundofile
setlocal undolevels=-123456
setlocal winfixheight
setlocal nowinfixwidth
set nowrap
setlocal nowrap
setlocal wrapmargin=0
silent! normal! zE
let s:l = 1 - ((0 * winheight(0) + 25) / 50)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
1
normal! 0
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 118 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 119 + 119) / 238)
tabnext 1
if exists('s:wipebuf')
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 shortmess=filnxtToOI
let s:sx = expand("<sfile>:p:r")."x.vim"
if file_readable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &so = s:so_save | let &siso = s:siso_save
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
