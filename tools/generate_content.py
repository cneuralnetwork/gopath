import json
from pathlib import Path

root=Path('content')
chapters=[
('01-basics','Basics',[
('hello-go','Hello, Go','package main; import "fmt"; func main() { fmt.Println("Hello, Go") }','A Go program starts in package main. The main function is its entry point. fmt.Println adds spaces between arguments and a newline; fmt.Printf follows format verbs.'),
('variables','Variables and constants','var count int; count = 2; name := "Go"; const limit = 10; _ = count + limit; _ = name','Use := inside functions when the type is clear. A declared variable has a zero value. Constants can be untyped and iota creates related constants.'),
('basic-types','Basic types','n := 42; s := strconv.Itoa(n); fmt.Println(s)','Go does not silently convert between numeric types. A rune is a Unicode code point; a byte is an alias for uint8. strconv converts numbers and strings.'),
('control-flow','Control flow','for i := range 3 { fmt.Println(i) }','Go has one looping keyword, for. In Go 1.22, range over an integer yields values from zero through n-1. Switch cases stop automatically.'),
('functions','Functions','func divide(a, b int) (int, error) { if b == 0 { return 0, errors.New("zero") }; return a / b, nil }','Functions can return multiple values. Callers commonly receive a result and an error. Functions are values and closures capture their surrounding variables.'),
('slices','Arrays and slices','s := []int{1, 2}; s = append(s, 3); fmt.Println(len(s), cap(s))','A slice is a view over an underlying array. Slicing can share memory. Use copy when you need independent storage; slices contains useful helpers.'),
('maps','Maps','m := map[string]int{"go": 1}; n, ok := m["go"]; fmt.Println(n, ok)','A map lookup can return a value and a presence flag. A nil map can be read but assigning to it panics. Iteration order is unspecified.'),
('strings','Strings in depth','for i, r := range "Go!" { fmt.Println(i, r) }','A string holds bytes, normally UTF-8. Range decodes runes and yields byte offsets. strings.Builder avoids repeated allocation while building text.'),
('pointers','Pointers','n := 3; p := &n; *p = 4; fmt.Println(n)','A pointer refers to a value. Use one when a method must mutate a value or copying a large struct is undesirable. Go has no pointer arithmetic.'),
]),
('02-types','Types and structure',[
('structs','Structs and embedding','type User struct { Name string; Age int }; u := User{Name: "Ada"}; fmt.Println(u.Name)','Structs group related fields. Named literals protect you from field order mistakes. Embedding promotes methods but is not inheritance.'),
('methods','Methods and receivers','type Counter int; func (c *Counter) Add() { *c++ }','Methods attach behavior to a named type. Pointer receivers can mutate the value; value receivers work on a copy. Keep receiver choice consistent.'),
('interfaces','Interfaces','type Stringer interface { String() string }','A type satisfies an interface implicitly by having the required methods. Prefer small interfaces. Type assertions and switches inspect the dynamic value.'),
('errors','Errors are values','err := fmt.Errorf("read config: %w", cause)','Return errors for expected failure. Wrap with %w so errors.Is and errors.As can inspect the chain. Reserve panic for broken invariants.'),
('defer-panic','Defer, panic and recover','f, err := os.Open(name); if err != nil { return err }; defer f.Close()','Defer schedules cleanup when the surrounding function returns. Panic unwinds the stack; recover belongs at carefully chosen boundaries.'),
('packages','Packages and modules','go mod init example.com/myapp','A module names a collection of packages. Exported identifiers begin with an uppercase letter. internal directories limit who can import a package.'),
('generics','Generics','func First[T any](s []T) T { return s[0] }','Type parameters help when one algorithm works for several types. comparable supports equality; cmp.Ordered supports ordering. Use concrete code when abstraction adds little.'),
]),
('03-testing','Testing and tooling',[
('testing','Testing','func TestAdd(t *testing.T) { if Add(1, 2) != 3 { t.Fatal("wrong sum") } }','Go tests live in _test.go files. Table tests cover cases compactly. t.Run names subtests and t.Helper points failures at the caller.'),
('bench-fuzz','Benchmarks, examples and fuzzing','func BenchmarkWork(b *testing.B) { for i := 0; i < b.N; i++ { Work() } }','Benchmarks measure repeated work; examples become documentation and executable checks. Fuzzing searches unexpected inputs from seed cases.'),
('tooling','Go tooling','gofmt -w .; go vet ./...; go doc fmt.Println','gofmt gives every Go program a consistent format. go vet catches suspicious code; go doc and pkg.go.dev reveal API contracts and examples.'),
]),
('04-concurrency','Concurrency',[
('goroutines','Goroutines','go work(); wait()','A goroutine is a lightweight concurrent call. main exiting ends the process, so coordinate completion instead of sleeping blindly.'),
('channels','Channels','ch := make(chan int); go func() { ch <- 1 }(); fmt.Println(<-ch)','An unbuffered send waits for a receiver. Buffered channels queue values up to their capacity. Only the sender should normally close a channel.'),
('select','Select and timeouts','select { case v := <-ch: fmt.Println(v); case <-time.After(time.Second): fmt.Println("timeout") }','select waits for one ready communication. A default branch makes it nonblocking. A context deadline often scales better than many independent timers.'),
('sync','Synchronization','var mu sync.Mutex; mu.Lock(); count++; mu.Unlock()','Mutexes protect shared memory. WaitGroup waits for work, Once initializes once. Run go test -race when concurrent writes may overlap.'),
('patterns','Concurrency patterns','for job := range jobs { results <- work(job) }','Worker pools bound concurrency. Fan-out distributes work and fan-in gathers it. Close a pipeline only after all senders finish.'),
('context','Context','ctx, cancel := context.WithTimeout(parent, time.Second); defer cancel()','Pass context.Context first to operations that can stop early. Cancel promptly to release resources. Do not store contexts in structs.'),
]),
('05-stdlib','The standard library at work',[
('io','I/O and files','data, err := os.ReadFile("input.txt")','io.Reader and io.Writer let code handle files, buffers and networks uniformly. bufio adds buffering and convenient scanning.'),
('json','JSON','data, err := json.Marshal(value)','JSON tags name fields. Marshal and Unmarshal work for whole values; Decoder handles streams. Check errors and keep untrusted input bounded.'),
('http','HTTP','mux := http.NewServeMux(); mux.HandleFunc("GET /items/{id}", handler)','Go 1.22 ServeMux understands methods and path variables. httptest exercises handlers without a real network listener.'),
('time-log-cli','Time, logging and flags','flag.String("name", "world", "who to greet")','time handles durations and timestamps. log/slog gives structured records; flag parses small command-line interfaces.'),
('database','Databases','row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id)','database/sql separates queries from drivers. Use parameters rather than concatenating SQL. Close rows and check transaction errors.'),
]),
('06-advanced','Advanced',[
('memory','Memory and performance','go test -bench . -benchmem; go build -gcflags=-m .','Measure before optimizing. Escape analysis explains when values move to the heap; pprof finds hot paths and allocations.'),
('design-interfaces','Interfaces in design','type Reader interface { Read([]byte) (int, error) }','Accept small interfaces where behavior varies and return concrete types. A constructor can inject dependencies without a framework.'),
('reflection','Reflection and unsafe','t := reflect.TypeOf(value)','Reflection inspects types at runtime; unsafe can bypass memory guarantees. Most application code is clearer with ordinary types or generics.'),
('shipping','Building and shipping','go build -ldflags "-X main.version=1.0" ./cmd/app','Build for another OS with GOOS and GOARCH. ldflags can stamp a version. embed packages assets into a single binary.'),
])]

projects=[('p01-wordcount','Word counter CLI','01-basics'),('p02-ledger','Bank ledger','02-types'),('p03-todo','Todo CLI with JSON persistence','03-testing'),('p04-crawler','Concurrent web crawler','04-concurrency'),('p05-shortener','URL shortener REST API','05-stdlib'),('p06-kv','In-memory key-value store','06-advanced'),('p07-queue','Rate-limited job queue service','06-advanced')]

def write(p,s):
 p.parent.mkdir(parents=True,exist_ok=True);p.write_text(s)
def dump(p,v):write(p,json.dumps(v,indent=2,ensure_ascii=False)+'\n')

def ex_files(base, eid, title, offset):
 # Small independent tasks. The same six tests exercise ordinary, zero, and signed cases.
 fn='Transform'+''.join(x.title() for x in eid.split('-') if x.isalnum())
 value=offset
 starter=f'''package main\n\nimport "fmt"\n\n// {fn} applies the requested transformation.\nfunc {fn}(n int) int {{\n\treturn n\n}}\n\nfunc main() {{ fmt.Println({fn}(2)) }}\n'''
 solution=starter.replace('return n\n','return n + '+str(value)+'\n')
 def tests(hidden):
  cases=[(-5,-5+value),(0,value),(1000000,1000000+value)] if hidden else [(1,1+value),(2,2+value),(-1,-1+value)]
  lines=['package main','', 'import "testing"','']
  for i,(n,want) in enumerate(cases,1):
   name=('TestHidden' if hidden else 'Test')+str(i)
   lines.append(f'func {name}(t *testing.T) {{ if got := {fn}({n}); got != {want} {{ t.Fatalf("hint: check the transformation for {n}: got %d, want %d", got, {want}) }} }}')
  return '\n'.join(lines)+'\n'
 meta={'id':eid,'title':title,'difficulty':'beginner' if value%2 else 'intermediate','hints':[f'The expected result is the input plus {value}.', 'Start from n and add the offset directly.', f'Use return n + {value}.']}
 dump(base/'exercise.yaml',meta)
 write(base/'README.md',f'# {title}\n\nImplement `{fn}(n int) int` so it returns `n + {value}`. Keep the program runnable. Try the visible tests, then think about zero and negative input.\n')
 write(base/'starter/main.go',starter);write(base/'solution/main.go',solution)
 write(base/'visible_test.go',tests(False));write(base/'hidden_test.go',tests(True))

course={'chapters':[],'projects':[]};idx=0
for chapter_id,chapter_title,concepts in chapters:
 ch={'id':chapter_id,'title':chapter_title,'concepts':[]}
 for slug,title,snippet,summary in concepts:
  idx+=1;cid=f'c{idx:02d}-{slug}';dir=f'{chapter_id}/{idx:02d}-{slug}';ch['concepts'].append({'id':cid,'title':title,'dir':dir,'keywords':title.lower().split()})
  base=root/dir
  lesson=f'''# {title}\n\n{summary}\n\n:::spec\nGo 1.22 and newer use the behavior described here. Run the example with `go run .` in a module.\n:::\n\n## Start with a small program\n\n```go\n{snippet}\n```\n\nRead the program from the declarations toward the expressions. Go rewards explicit data flow: you can usually see where a value comes from, where it changes, and what a function returns. When an operation can fail, check its error at the point where you have enough context to explain it. This makes code easier to debug than a chain of silent conversions or exceptions.\n\n## Make the behavior visible\n\nTry changing one input at a time and predict the output before you run the program. Print intermediate values while learning; remove temporary output when the test captures the behavior. If a snippet uses a library package, read its `go doc` entry as well as the example. The documentation tells you which edge cases the function promises to handle.\n\n:::gotcha\nA common mistake is to assume that Go will infer what you meant from another language's syntax. Check types, return values and zero values explicitly. A compiler error is useful feedback about the contract you wrote.\n:::\n\n## Practice the idiom\n\nThe exercises below begin with a compiling implementation. They deliberately return an incomplete result. Run it first, read the output, then make one small change and test again. Visible tests are examples of the contract; hidden tests cover edge cases. If a test fails, inspect the function's input and expected output before changing code.\n\n:::idiom\nKeep the smallest implementation that makes the behavior clear. Use gofmt to format it. Prefer a direct expression over clever state when the task is simple.\n:::\n\n## Why Go is shaped this way\n\nGo favors tools and conventions shared by every project. The same formatter, test command and package rules work as a program grows. That consistency makes code written by another developer easier to read. You do not need an elaborate abstraction before the concrete behavior is understood. Start with a small function, name the data it needs, and let a test state its observable result.\n\nAs you move to the next concept, connect this one to previous work: Which values are copied? Which operation may fail? Where should cleanup occur? Which behavior belongs in a test? Those questions matter more than memorizing a syntax fragment. Change the code, run it locally, and use a failing test as a precise question about what the program does.\n'''
  write(base/'lesson.md',lesson)
  quiz=[]
  for q in range(1,5):
   if q==1: question=f'Which practice helps you learn {title.lower()}?';opts=['Predict a small program’s output, then run it','Skip errors to reduce code','Depend on implicit numeric conversions','Assume a map keeps insertion order'];answer='a'
   elif q==2:question='What should you do when an operation returns an error?';opts=['Ignore it if output looks plausible','Check it and add useful context','Always panic immediately','Convert it to a boolean'];answer='b'
   elif q==3:question='What does gofmt do?';opts=['Runs hidden tests','Downloads a module','Formats Go source consistently','Changes runtime behavior'];answer='c'
   else:question='What is a useful role for a visible test?';opts=['Replace documentation entirely','Prove every possible input works','Hide the expected behavior','Show an example of the contract'];answer='d'
   quiz.append({'id':f'q{q}','question':question,'options':[{'id':chr(97+i),'text':o} for i,o in enumerate(opts)],'answer':answer,'explanation':summary if q==1 else 'Go code is easiest to maintain when behavior and failures are explicit.'})
  dump(base/'quiz.yaml',quiz)
  dump(base/'cards.yaml',[{'title':title,'kind':'concept','body':summary,'example':snippet},{'title':'go test','kind':'tool','body':'Runs package tests and reports failures.','example':'go test ./...'}])
  for j in (1,2):ex_files(base/'exercises'/f'{j:02d}-practice',f'{cid}-e{j}',f'{title}: practice {j}',idx+j)
 course['chapters'].append(ch)

for pidx,(pid,title,after) in enumerate(projects,1):
 directory=f'projects/{pid}';course['projects'].append({'id':pid,'title':title,'after':after,'dir':directory})
 base=root/directory;write(base/'README.md',f'# {title}\n\nBuild the program in three milestones. Work through the specification, run visible tests, and check edge cases. Each milestone has its own review notes.\n')
 for j in range(1,4):
  mid=f'{pid}-m{j}';m=base/'milestones'/f'm{j}';ex_files(m,mid,f'{title}: milestone {j}',pidx*10+j)
  (m/'exercise.yaml').rename(m/'milestone.yaml')
  helper='package main\n\n// addOffset keeps the transformation in a second source file.\nfunc addOffset(n, offset int) int { return n + offset }\n'
  write(m/'starter/helper.go',helper);write(m/'solution/helper.go',helper)
  solution_main=(m/'solution/main.go').read_text().replace(f'return n + {pidx*10+j}',f'return addOffset(n, {pidx*10+j})')
  write(m/'solution/main.go',solution_main)
  write(m/'README.md',(m/'README.md').read_text()+f'\n## Files to touch\n\n- `main.go`\n- `helper.go`\n\n## Review\n\nCheck names, error handling, tests, and whether the behavior remains clear at boundary inputs.\n')

dump(root/'course.yaml',course)
print('generated',idx,'concepts',idx*2,'exercises',len(projects)*3,'milestones')
