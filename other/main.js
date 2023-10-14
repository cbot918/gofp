const log = console.log

const from = (v) => ({
  then: fn => from(fn(v)),
  return: () => v
})


const answer = from(3)
                .then(x=>x+1)
                .then(x=>x*5)
                .return()

log(answer)
